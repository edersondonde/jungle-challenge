package handler

import (
	"errors"
	"net/http"

	"github.com/edersondonde/jungle-challenge/controller"
	"github.com/edersondonde/jungle-challenge/utils"
	"github.com/gin-gonic/gin"
)

// GetClients serves /info requests
func GetClients(c *gin.Context) {
	id := c.Query("clientId")

	if id != "" {
		getClientById(c, id)
	}

	startBirthDate := c.Query("startBirthDate")
	endBirthDate := c.Query("endBirthDate")

	getClientsBetweenBirthDates(c, startBirthDate, endBirthDate)
}

// getClientById handler to get the user by ID
func getClientById(c *gin.Context, id string) {
	client, err := controller.GetClientById(id)
	if err != nil {
		if errors.Is(err, utils.ErrClientNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, client)
}

// getClientsBetweenBirthDates handler to get users that have birthDate between the requested dates
func getClientsBetweenBirthDates(c *gin.Context, startBirthDate string, endBirthDate string) {
	clients, err := controller.GetClientsBetweenBirthDates(startBirthDate, endBirthDate)
	if err != nil {
		if errors.Is(err, utils.ErrStartDateWrongFormat) ||
		errors.Is(err, utils.ErrEndDateWrongFormat)  ||
		errors.Is(err, utils.ErrEndDateBeforeStartDate){
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, clients)
}

// getClientsBetweenBirthDates handler to search for users by name
func SearchClientByName(c *gin.Context) {
	name := c.Query("name")

	client, err := controller.SearchClientByName(name)

	if err != nil {
		if errors.Is(err, utils.ErrClientNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, client)
}
