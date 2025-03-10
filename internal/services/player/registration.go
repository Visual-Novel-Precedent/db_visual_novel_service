package player

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"errors"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

const (
	PlayerNotFoundError  = "player data not found"
	DefaultSoundSettings = 50
)

func Registration(email string, name string, password string, db *gorm.DB) (int64, error) {
	player, err := storage.SelectPlayerWIthEmail(db, email)

	if player.Email == email {
		return 0, errors.New("player with this email is already exist")
	}

	id := generateUniqueId()

	newPlayer := models.Player{
		Id:                id,
		Email:             email,
		Name:              name,
		Password:          password,
		CompletedChapters: []int64{},
		ChaptersProgress:  map[int64]int64{},
		SoundSettings:     DefaultSoundSettings,
	}

	_, err = storage.RegisterPlayer(db, newPlayer)

	log.Println(err)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func generateUniqueId() int64 {
	// Получаем текущее время в миллисекундах (48 бит)
	timestamp := time.Now().UnixMilli()

	// Генерируем 16 случайных бит
	random := rand.Int31n(1 << 16)

	// Объединяем timestamp и random в 64-битное число
	return (int64(timestamp) << 16) | int64(random)
}
