package chapter

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
)

func GetChaptersByUserId(db *gorm.DB, id int64) ([]models.Chapter, error) {

	_, err := storage.SelectAdminWIthId(db, id)

	if err == nil {
		chapters, err := storage.GetChaptersForAdmin(db)

		if err != nil {
			return nil, err
		}

		return chapters, nil
	}

	_, err = storage.SelectPlayerWIthId(db, id)

	if err != nil {
		return nil, err
	}

	chapters, err := storage.FindPublishedChapters(db)

	if err != nil {
		return nil, err
	}

	return chapters, nil
}
