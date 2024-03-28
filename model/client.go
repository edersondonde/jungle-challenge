package model

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id         uuid.UUID `json:"id"`
	Sex        string    `json:"sex"`
	DayOfBirth time.Time `json:"dayOfBirth"`
	Name       string    `json:"name"`
}
