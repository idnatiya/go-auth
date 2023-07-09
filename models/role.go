package models

import "time"

type Role struct {
	ID         uint         `gorm:"primary_key"`
	RoleName   string       `gorm:"unique"`
	Permission []Permission `gorm:"many2many:role_permissions;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
