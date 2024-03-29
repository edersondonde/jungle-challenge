package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/edersondonde/jungle-challenge/controller"
	"github.com/edersondonde/jungle-challenge/domain"
	"github.com/edersondonde/jungle-challenge/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {

	initConfig()

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

	url := viper.GetString("server.url")
	port := viper.GetString("server.port")
	address := fmt.Sprintf("%v:%v", url, port)

	router.Run(address)
}

func initDatabase() *sql.DB {

	url := viper.GetString("database.url")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	pass := viper.GetString("database.password")
	dbName := viper.GetString("database.name")

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, pass, url, port, dbName)

	fmt.Println(connStr)

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
