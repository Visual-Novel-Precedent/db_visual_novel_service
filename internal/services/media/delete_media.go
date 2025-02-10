package media

import (
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
)

func DeleteMedia(id int64, db *gorm.DB) (int64, error) {

	_, err := storage.DeleteMedia(db, id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
