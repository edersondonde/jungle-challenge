package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

const help = "it is expected one of the following commands:\n" +
	"    migrate up\n" +
	"    migrate down\n" +
	"    loadDB <FILE_PATH>\n" +
	"    init <FILEPATH>."

func main() {

	args := os.Args[1:]

	if len(args) != 2 {
		log.Fatal(help)
	}

	initConfig()

	if args[0] == "migrate" {
		migrateDB(args[1])
	} else if args[0] == "loadDB" {
		loadDB(args[1])
	} else if args[0] == "init" {
		migrateDBStep()
		fmt.Println("First migration step done, loading DB")
		loadDB(args[1])
		fmt.Println("DB loading, starting second migration step")
		migrateDBStep()
		fmt.Println("Database initialization finished")
	}

}

func migrateDB(command string) {
	if command != "up" && command != "down" {
		log.Fatal(help)
	}

	connStr := getDBConnectionUrl()

	m, err := migrate.New(
		"file://jungle",
		connStr)
	if err != nil {
		log.Fatal(err)
	}
	if command == "up" {
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("migration up done")
	} else if command == "down" {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("migration down done")
	}
}

func migrateDBStep() {
	connStr := getDBConnectionUrl()

	m, err := migrate.New(
		"file://jungle",
		connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Steps(1); err != nil {
		log.Fatal(err)
	}
	fmt.Println("migration step done")
}

func loadDB(filePath string) {
	url := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	pass := viper.GetString("database.password")
	dbName := viper.GetString("database.name")

	command := "psql"
	args := fmt.Sprintf("%v -U %v -p %v -h %v -c \"\\copy client from '%v' with (DELIMITER ',', format csv, header)\"", dbName, user, port, url, filePath)

	cmd := exec.Command(command, args)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("PGPASSWORD=%v", pass))

	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("could not run command: %v", err)
	}
	fmt.Println("Output: ", string(out))

}

func getDBConnectionUrl() string {
	url := viper.GetString("database.url")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	pass := viper.GetString("database.password")
	dbName := viper.GetString("database.name")

	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, pass, url, port, dbName)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("ederson")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
