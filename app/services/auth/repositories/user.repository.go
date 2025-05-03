package repositories

import (
	"app/services/auth/entities"

	"gorm.io/gorm"
)

type UserRepository struct {}

func (r *UserRepository) FindByUsername(db *gorm.DB, user *entities.User, username string) error {
	return db.Where("username = ?", username).Take(user).Error
}