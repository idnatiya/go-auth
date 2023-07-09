package models

import "time"

type Permission struct {
	ID             uint   `gorm:"primary_key"`
	PermissionName string `gorm:"unique"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
