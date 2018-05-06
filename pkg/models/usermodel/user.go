package usermodel

import "time"

type User struct {
	ID         uint64    `json:"id"`
	LastUpdate time.Time `json:"last_update"`
}
