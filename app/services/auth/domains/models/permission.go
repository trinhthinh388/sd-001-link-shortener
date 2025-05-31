package models

import "time"

type Permission struct {
	Permission string    `gorm:"type:varchar;size:255;primaryKey"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	Roles []*Role `gorm:"many2many:roles_permissions;foreignKey:Permission"`
}
