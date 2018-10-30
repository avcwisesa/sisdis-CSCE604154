package model

import (
	"time"
)

type Customer struct {
	ID                uint       `json:"id" gorm:"primary_key"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
	Domisili          string     `json:"domisili"`
	UserID            uint       `json:"user_id" gorm:"UNIQUE"`
	Balance           uint       `json:"balance"`
}