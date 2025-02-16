package media

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func CreateMedia(
	file []byte,
	contentType string,
	db *gorm.DB,
) (int64, error) {
	id := generateUniqueId()

	newMedia := models.Media{
		Id:          id,
		FileData:    file,
		ContentType: contentType,
	}

	_, err := storage.RegisterMedia(db, newMedia)

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
