package domain

import (
	"database/sql"
	"time"

	"github.com/edersondonde/jungle-challenge/model"
	"github.com/edersondonde/jungle-challenge/utils"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) ClientRepository {
	return ClientRepository{
		db: db,
	}
}

func (c ClientRepository) FindById(id string) (*model.Client, error) {
	query := "SELECT uid, birthday, sex, name FROM client WHERE uid = $1"
	client := &model.Client{}

	err := c.db.QueryRow(query, id).Scan(&client.Id, &client.Birthday, &client.Sex, &client.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrClientNotFound
		}
		return nil, err
	}
	return client, nil
}

func (c ClientRepository) FindByName(name string) (*model.Client, error) {
	query := "SELECT uid, birthday, sex, name FROM client WHERE name LIKE $1 LIMIT 1"
	client := &model.Client{}

	err := c.db.QueryRow(query, name+"%").Scan(&client.Id, &client.Birthday, &client.Sex, &client.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrClientNotFound
		}
		return nil, err
	}
	return client, nil
}

func (c ClientRepository) GetClientsBetweenBirthdays(startBirthday time.Time, endBirthday time.Time) ([]*model.Client, error) {
	query := "SELECT uid, birthday, sex, name FROM client WHERE birthday BETWEEN $1 AND $2"
	
	rows, err := c.db.Query(query, startBirthday, endBirthday)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrClientNotFound
		}
		return nil, err
	}

	defer rows.Close()

	result := []*model.Client{}

	for rows.Next() {
		client := &model.Client{}
		err := rows.Scan(&client.Id, &client.Birthday, &client.Sex, &client.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, client)
	}

	return result, nil
}
