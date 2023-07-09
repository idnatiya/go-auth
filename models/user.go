package models

import "time"

type User struct {
	ID        uint   `gorm:"primary_key"`
	FirstName string `gorm:"index"`
	LastName  string `gorm:"index"`
	UserName  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
