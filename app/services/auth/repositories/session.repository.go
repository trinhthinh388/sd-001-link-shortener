package repositories

import (
	"gorm.io/gorm"
)

type SessionRepository struct {}

func (r *SessionRepository) UpdateByUserId(db *gorm.DB, token string, userId string) error {
	return db.Exec("INSERT INTO sessions (id, user_id) VALUES(?, ?) ON DUPLICATE KEY UPDATE id = ?;", token, userId, token).Error
}