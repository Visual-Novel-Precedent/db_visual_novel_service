package request

import (
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
)

const (
	RejectedRequestStatus = -1
)

func RejectRequest(id int64, db *gorm.DB) error {
	request, err := storage.SelectRequestWIthId(db, id)

	if err != nil {
		return err
	}

	request.Status = RejectedRequestStatus

	_, err = storage.UpdateRequest(db, request.Id, request)

	_, err = storage.DeleteRequest(db, id)

	return err
}
