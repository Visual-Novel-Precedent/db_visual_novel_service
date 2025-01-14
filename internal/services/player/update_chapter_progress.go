package player

import (
	"db_novel_service/internal/storage"
	"errors"
	"gorm.io/gorm"
	"slices"
)

const (
	EndNodeId = -1
)

func UpdateChapterProgress(id int64, chapterId int64, nodeId int64, endFlag bool, db *gorm.DB) error {
	user, err := storage.SelectPlayerWIthId(db, id)

	if err != nil {
		return err
	}

	if user.ChaptersProgress == nil {
		return errors.New("error to get player progress")
	}

	if endFlag {
		user.ChaptersProgress[chapterId] = EndNodeId

		_, err = storage.UpdatePlayer(db, id, user)

		if err != nil {
			return err
		}

		return nil
	}

	chapter, err := storage.SelectChapterWIthId(db, chapterId)

	if err != nil {
		return err
	}

	if chapter.Nodes == nil || !slices.Contains(chapter.Nodes, nodeId) {
		return errors.New("chapter not contains this node")
	}

	user.ChaptersProgress[chapterId] = nodeId

	_, err = storage.UpdatePlayer(db, id, user)

	if err != nil {
		return err
	}

	return nil
}
