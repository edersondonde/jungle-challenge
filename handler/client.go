package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/edersondonde/jungle-challenge/controller"
	"github.com/edersondonde/jungle-challenge/utils"
	"github.com/gin-gonic/gin"
)

type Client struct {
	controller controller.Client
}

func NewClientHandler(controller controller.Client) Client {
	return Client{controller: controller}
}

// GetClients serves /info requests
func (h Client) GetClients(c *gin.Context) {
	id := c.Query("clientId")

	if id != "" {
		h.getClientById(c, id)
		return
	}

	startBirthday := c.Query("startBirthday")
	endBirthday := c.Query("endBirthday")

	h.getClientsBetweenBirthdays(c, startBirthday, endBirthday)
}

// getClientById handler to get the user by ID
func (h Client) getClientById(c *gin.Context, id string) {
	client, err := h.controller.GetClientById(id)
	if err != nil {

		fmt.Println("Error searching for Client by ID")
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
func (h Client) getClientsBetweenBirthdays(c *gin.Context, startBirthday string, endBirthday string) {
	clients, err := h.controller.GetClientsBetweenBirthdays(startBirthday, endBirthday)
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
func (h Client) SearchClientByName(c *gin.Context) {
	name := c.Query("name")

	client, err := h.controller.SearchClientByName(name)

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
