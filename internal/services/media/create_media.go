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
	// Получаем текущее время в миллисекундах (48 бит)
	timestamp := time.Now().UnixMilli()

	// Генерируем 16 случайных бит
	random := rand.Int31n(1 << 16)

	// Объединяем timestamp и random в 64-битное число
	return (int64(timestamp) << 16) | int64(random)
}
