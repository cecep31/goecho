package models

import "time"

// USers
type Users struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"Username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
