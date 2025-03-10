package request

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"log"
)

func GetMyRequests(id int64, db *gorm.DB) ([]models.Request, error) {
	admin, err := storage.SelectAdminWithId(db, id)

	log.Println("admin", admin.AdminStatus)

	if err != nil {
		return nil, err
	}

	if admin.AdminStatus != 1 {
		return nil, nil
	}

	requests, err := storage.GetAllRequests(db)

	return requests, err
}
