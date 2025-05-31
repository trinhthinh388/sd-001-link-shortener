package models

import "time"

type Role struct {
	Role      string    `gorm:"type:varchar;size:255;primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Permission []*Permission `gorm:"many2many:roles_permissions;foreignKey:Role"`
}
