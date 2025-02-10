package media

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
)

func GetMediaById(id int64, db *gorm.DB) (*models.Media, error) {

	media, err := storage.SelectMediaWIthId(db, id)

	if err != nil {
		return nil, err
	}

	return &media, nil
}
