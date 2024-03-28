package controller

import (
	"time"

	"github.com/edersondonde/jungle-challenge/domain"
	"github.com/edersondonde/jungle-challenge/model"
	"github.com/edersondonde/jungle-challenge/utils"
)

type Client struct {
	repository domain.ClientRepository
}

func NewClientController(clientRepository domain.ClientRepository) Client {
	return Client{
		repository: clientRepository,
	}
}

// GetClientById search for a user by ID
func (c Client) GetClientById(id string) (*model.Client, error) {
	client, err := c.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// GetClientsBetweenBirthdays retrieves a list of users that have birthdays between requested dates
func (c Client) GetClientsBetweenBirthdays(startBirthdayStr string, endBirthdayStr string) ([]*model.Client, error) {
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

	resultClients, err := c.repository.GetClientsBetweenBirthdays(startBirthday, endBirthday)
	if err != nil {
		return nil, err
	}

	return resultClients, nil
}

// SearchClientByName finds the first user that matches the requested name
func (c Client) SearchClientByName(name string) (*model.Client, error) {

	client, err := c.repository.FindByName(name)
	if err != nil {
		return nil, err
	}

	return client, nil
}
