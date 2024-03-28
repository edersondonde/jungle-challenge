package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/edersondonde/jungle-challenge/controller"
	"github.com/edersondonde/jungle-challenge/domain"
	"github.com/edersondonde/jungle-challenge/handler"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {
	db := initDatabase()

	defer db.Close()

	clientRepository := initRepositories(db)
	initServer(clientRepository)
}

func initServer(clientRepository domain.ClientRepository) {

	controller := controller.NewClientController(clientRepository)
	clientHandler := handler.NewClientHandler(controller)

	router := gin.Default()
	router.GET("/info/", clientHandler.GetClients)
	router.GET("/search/", clientHandler.SearchClientByName)

	router.Run("localhost:8080")
}

func initDatabase() *sql.DB {
	connStr := "postgres://postgres:pass@localhost:5432/jungle-challenge?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("The database is connected")

	return db
}

func initRepositories(db *sql.DB) domain.ClientRepository {
	repository := domain.NewClientRepository(db)
	return repository
}