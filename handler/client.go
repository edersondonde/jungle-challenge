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

	startBirthday := c.Query("startBirthday")
	endBirthday := c.Query("endBirthday")

	getClientsBetweenBirthdays(c, startBirthday, endBirthday)
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

// getClientsBetweenBirthdays handler to get users that have birthday between the requested dates
func getClientsBetweenBirthdays(c *gin.Context, startBirthday string, endBirthday string) {
	clients, err := controller.GetClientsBetweenBirthdays(startBirthday, endBirthday)
	if err != nil {
		if errors.Is(err, utils.ErrStartDateWrongFormat) ||
			errors.Is(err, utils.ErrEndDateWrongFormat) ||
			errors.Is(err, utils.ErrEndDateBeforeStartDate) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, clients)
}

// SearchClientByName handler to search for users by name
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
