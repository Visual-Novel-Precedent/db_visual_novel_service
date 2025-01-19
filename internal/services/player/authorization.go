package player

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"errors"
	"gorm.io/gorm"
)

func Authorization(email string, password string, db *gorm.DB) (*models.Player, error) {
	user, err := storage.SelectPlayerWIthEmail(db, email)

	if err != nil {
		return nil, errors.New("error to get user from storage")
	}

	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return &user, err
}
