package controller

import (
	"strings"
	"time"

	"github.com/edersondonde/jungle-challenge/model"
	"github.com/edersondonde/jungle-challenge/utils"
	"github.com/google/uuid"
)

// GetClientById search for a user by ID
func GetClientById(id string) (*model.Client, error) {
	clients := getClients()

	for _, client := range clients {
		if client.Id.String() == id {
			return client, nil
		}
	}

	return nil, utils.ErrClientNotFound
}

// GetClientsBetweenBirthDates retrieves a list of users that have birthDates between requested dates
func GetClientsBetweenBirthDates(startBirthDateStr string, endBirthDateStr string) ([]*model.Client, error) {
	clients := getClients()
	startBirthDate, err := time.Parse("2006-01-02", startBirthDateStr)
	if err != nil {
		return nil, utils.ErrStartDateWrongFormat
	}
	endBirthDate, err := time.Parse("2006-01-02", endBirthDateStr)
	if err != nil {
		return nil, utils.ErrEndDateWrongFormat
	}
	if endBirthDate.Before(startBirthDate) {
		return nil, utils.ErrEndDateBeforeStartDate
	}

	var resultClients []*model.Client

	for _, client := range clients {
		if client.DayOfBirth.After(startBirthDate) && client.DayOfBirth.Before(endBirthDate) {
			resultClients = append(resultClients, client)
		}
	}

	return resultClients, nil
}

// SearchClientByName finds the first user that matches the requested name
func SearchClientByName(name string) (*model.Client, error) {
	clients := getClients()

	for _, client := range clients {
		if strings.HasPrefix(client.Name, name) {
			return client, nil
		}
	}

	return nil, utils.ErrClientNotFound
}

// Temporary while database is not implemented
func getClients() []*model.Client {
	clients := []*model.Client{
		{Id: uuid.MustParse("3acace50-c45c-4e30-af42-156e3335d0d7"), Sex: "male", DayOfBirth: time.Now().AddDate(-5, 0, 0), Name: "Ederson"},
		{Id: uuid.New(), Sex: "female", DayOfBirth: time.Now().AddDate(-7, 0, 0), Name: "Ana Flavia"},
	}
	return clients
}
