package chapter

import (
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"time"
)

func UpdateChapter(
	id int64,
	name string,
	nodes []int64,
	characters []int64,
	updateAuthorId int64,
	db *gorm.DB,
) error {
	chapter, err := storage.SelectChapterWIthId(db, id)

	if err != nil {
		return err
	}

	newChapter := chapter

	newChapter.UpdatedAt[time.Now()] = updateAuthorId

	if name != "" {
		newChapter.Name = name
	}

	if nodes != nil {
		newChapter.Nodes = nodes
	}

	if characters != nil {
		newChapter.Characters = characters
	}

	_, err = storage.UpdateChapter(db, id, newChapter)

	return err
}
