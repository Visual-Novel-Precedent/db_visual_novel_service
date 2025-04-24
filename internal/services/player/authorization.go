package player

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"errors"
	"gorm.io/gorm"
	"log"
)

func Authorization(email string, password string, db *gorm.DB) (*models.Player, error) {
	user, err := storage.SelectPlayerWIthEmail(db, email)

	log.Println("ошибка", err)

	log.Println("найденный игрок", user)

	if err != nil {
		return nil, errors.New("error to get user from storage")
	}

	log.Println(user.Password, "666")
	log.Println("888", password)

	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	log.Println("найденный юзер", user)

	return &user, err
}
