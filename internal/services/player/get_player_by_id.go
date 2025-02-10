package player

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"errors"
	"gorm.io/gorm"
)

func GetPlayerById(id int64, db *gorm.DB) (*models.Player, error) {
	user, err := storage.SelectPlayerWIthId(db, id)

	if err != nil {
		return nil, errors.New("error to get user from storage")
	}

	return &user, err
}
