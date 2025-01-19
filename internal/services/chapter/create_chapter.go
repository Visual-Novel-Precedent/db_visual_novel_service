package chapter

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

const (
	DefaultStatus = 0
)

func CreateDefaultChapter(authorId int64, db *gorm.DB) (int64, error) {

	id := generateUniqueId()

	newChapter := models.Chapter{
		Id:        id,
		Author:    authorId,
		Status:    DefaultStatus,
		UpdatedAt: map[time.Time]int64{time.Now(): id},
	}

	_, err := storage.RegisterChapter(db, newChapter)

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
