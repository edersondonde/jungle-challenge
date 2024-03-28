package controller

import (
	"strings"
	"time"

	"github.com/edersondonde/jungle-challenge/model"
	"github.com/edersondonde/jungle-challenge/utils"
)

// GetClientById search for a user by ID
func GetClientById(id string) (*model.Client, error) {
	clients := getClients()

	for _, client := range clients {
		if client.Id == id {
			return client, nil
		}
	}

	return nil, utils.ErrClientNotFound
}

// GetClientsBetweenBirthdays retrieves a list of users that have birthdays between requested dates
func GetClientsBetweenBirthdays(startBirthdayStr string, endBirthdayStr string) ([]*model.Client, error) {
	clients := getClients()
	startBirthday, err := time.Parse("2006-01-02", startBirthdayStr)
	if err != nil {
		return nil, utils.ErrStartDateWrongFormat
	}
	endBirthday, err := time.Parse("2006-01-02", endBirthdayStr)
	if err != nil {
		return nil, utils.ErrEndDateWrongFormat
	}
	if endBirthday.Before(startBirthday) {
		return nil, utils.ErrEndDateBeforeStartDate
	}

	var resultClients []*model.Client

	for _, client := range clients {
		if client.Birthday.After(startBirthday) && client.Birthday.Before(endBirthday) {
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
		{Id: "000001", Sex: "male", Birthday: time.Now().AddDate(-5, 0, 0), Name: "Ederson"},
		{Id: "000002", Sex: "female", Birthday: time.Now().AddDate(-7, 0, 0), Name: "Ana Flavia"},
	}
	return clients
}
