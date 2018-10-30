package model

import (
	"time"
)

type Customer struct {
	ID                uint       `json:"id" gorm:"primary_key"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
	UserID            uint       `json:"user_id" gorm:"UNIQUE"`
	Domisili          string     `json:"domisili"`
	Name              string     `json:"name"`
	Balance           uint       `json:"balance"`
}