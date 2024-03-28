package model

import (
	"time"
)

type Client struct {
	Id       string    `json:"id"`
	Sex      string    `json:"sex"`
	Birthday time.Time `json:"birthday"`
	Name     string    `json:"name"`
}
