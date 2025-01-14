package player

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"errors"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

const (
	PlayerNotFoundError = "admin data not found"
)

func Registration(email string, name string, password string, db *gorm.DB) (int64, error) {
	_, err := storage.SelectPlayerWIthEmail(db, email)

	if err == nil {
		return 0, errors.New("player with this email is already exist")
	}

	if err.Error() != PlayerNotFoundError {
		return 0, err
	}

	id := generateUniqueId()

	newPlayer := models.Player{
		Id:       id,
		Email:    email,
		Password: password,
	}

	_, err = storage.RegisterPlayer(db, newPlayer)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func generateUniqueId() int64 {
	now := time.Now().UnixNano()
	random := rand.Int63n(1 << 32) // 32-битное случайное число
	return now ^ random
}
