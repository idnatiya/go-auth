package models

import "time"

type Permission struct {
	ID             uint   `gorm:"primary_key"`
	PermissionName string `gorm:"unique"`
	Role           []Role `gorm:"many2many:role_permissions;"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
