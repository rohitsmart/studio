package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique"`
	MobileNumber string `json:"mobile_number" gorm:"unique"`
	Password     string `json:"-"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Role         string `json:"role" gorm:"not null"` // "customer", "freelancer", "studio"
}
