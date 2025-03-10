package request

import (
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"log"
)

const (
	RejectedRequestStatus = -1
)

func RejectRequest(id int64, db *gorm.DB) error {
	request, err := storage.SelectRequestWithId(db, id)

	if err != nil {
		return err
	}

	request.Status = RejectedRequestStatus

	_, err = storage.DeleteRequest(db, id)

	if request.Type == 1 {
		chapter, err := storage.SelectChapterWIthId(db, request.RequestedChapterId)

		if err != nil {
			return err
		}

		chapter.Status = 1

		_, err = storage.UpdateChapter(db, chapter.Id, chapter)

		if err != nil {
			return err
		}
	}

	log.Println("ошибка удаления запооса", err)

	return err
}
