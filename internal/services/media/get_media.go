package media

import (
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
)

func GetMedia(db *gorm.DB) ([]int64, error) {

	media, err := storage.SelectMedia(db)

	if err != nil {
		return nil, err
	}

	var ids []int64

	for _, m := range media {
		ids = append(ids, m.Id)
	}

	return ids, nil
}
