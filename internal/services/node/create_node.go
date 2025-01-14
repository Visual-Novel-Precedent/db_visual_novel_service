package node

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func CreateNode(chapterId int64, slug string, db *gorm.DB) (int64, error) {

	id := generateUniqueId()

	newNode := models.Node{
		Id:        id,
		Slug:      slug,
		ChapterId: chapterId,
	}

	_, err := storage.RegisterNode(db, newNode)

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
