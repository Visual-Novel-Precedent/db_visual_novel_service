package character

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func CreateCharacter(name string, slug string, db *gorm.DB) (int64, error) {

	id := generateUniqueId()

	newCharacter := models.Character{
		Id:   id,
		Name: name,
		Slug: slug,
	}

	_, err := storage.RegisterCharacter(db, newCharacter)

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
