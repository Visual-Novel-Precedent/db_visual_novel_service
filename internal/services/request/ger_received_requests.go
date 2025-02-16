package request

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"log"
)

func GetReceivedRequests(id int64, db *gorm.DB) ([]models.Request, error) {
	admin, err := storage.SelectAdminWithId(db, id)

	if err != nil {
		return nil, err
	}

	requestID := admin.RequestsReceived

	var requests []models.Request

	for _, i := range requestID {
		request, err := storage.SelectRequestWIthId(db, i)

		if err == nil {
			requests = append(requests, request)
		}
	}

	log.Println(requests)

	return requests, nil
}
