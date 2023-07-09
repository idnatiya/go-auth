package models

import "time"

type File struct {
	ID           uint `gorm:"primary_key"`
	FileName     string
	PathLocation string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
