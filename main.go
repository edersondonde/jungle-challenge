package main

import (
	"github.com/edersondonde/jungle-challenge/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	initServer()
}

func initServer() {
	router := gin.Default()
	router.GET("/info/", handler.GetClients)
	router.GET("/search/", handler.SearchClientByName)

	router.Run("localhost:8080")
}


