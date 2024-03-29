package main

import (
	"database/sql"
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
	"    loadDB <FILE_PATH>."

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
	} else {
		log.Fatal(help)
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

func loadDB(filePath string) {
	url := viper.GetString("database.url")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	pass := viper.GetString("database.password")
	dbName := viper.GetString("database.name")

	os.Setenv("PGPASSWORD", pass)

	defer os.Clearenv("PGPASSWORD")

	command := fmt.Sprintf("psql %v -U %v -p %v -h %v -c \"\\copy client from '%v' with (DELIMITER ',', format csv, header)", dbName, user, port, url, filePath)

	cmd := exec.Command(command)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	fmt.Println("Output: ", string(out))

}

func getDBConnectionUrl() string () {
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