package models

import (
	"time"
)

type Artikel struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title" gorm:"type:varchar(50)"`
	Content   string    `gorm:"type:text" json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
