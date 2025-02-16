package request

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

func CreateRequest(requestingAdminId int64, typeRequest int, requestedChapterId int64, db *gorm.DB) (int64, error) {

	id := generateUniqueId()

	log.Println(id)

	newRequest := models.Request{
		Id:                 id,
		Type:               typeRequest,
		RequestingAdmin:    requestingAdminId,
		RequestedChapterId: requestedChapterId,
	}

	_, err := storage.RegisterRequest(db, newRequest)

	if err != nil {
		return 0, err
	}

	admin, err := storage.SelectAdminWithId(db, requestingAdminId)

	if err != nil {
		return 0, err
	}

	admin.RequestSent = append(admin.RequestSent, id)

	_, err = storage.UpdateAdmin(db, admin.Id, admin)

	if err != nil {
		return 0, err
	}

	admins, err := storage.SelectAllSupeAdmins(db)

	for _, admin := range admins {
		ad, err := storage.SelectAdminWithId(db, admin.Id)

		if err == nil {
			ad.RequestsReceived = append(ad.RequestsReceived, id)
		}

		_, _ = storage.UpdateAdmin(db, admin.Id, ad)
	}

	return id, nil
}

func generateUniqueId() int64 {
	now := time.Now().UnixNano()
	random := rand.Int63n(1 << 32) // 32-битное случайное число
	return now ^ random
}
