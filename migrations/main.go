package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	m, err := migrate.New(
		"file://jungle",
		"postgres://postgres:pass@localhost:5432/jungle-challenge?sslmode=disable")
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
	connStr := "postgres://postgres:pass@localhost:5432/jungle-challenge?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("The database is connected")

	sql := fmt.Sprintf("COPY client "+
		"FROM '%v' "+
		"DELIMITER ',' "+
		"CSV HEADER;", filePath)

	result, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(result.RowsAffected())

}
