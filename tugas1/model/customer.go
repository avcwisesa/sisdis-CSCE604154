package model

import (
	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	UserID            string     `json:"user_id" gorm:"unique;type:varchar(20)"`
	Domisili          string     `json:"domisili" gorm:"type:varchar(30)"`
	Name              string     `json:"name" gorm:"type:varchar(40)"`
	Balance           uint       `json:"balance"`
}
