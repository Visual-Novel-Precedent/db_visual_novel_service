package request

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func CreateRequest(requestingAdminId int64, typeRequest int, requestedChapterId int64, db *gorm.DB) (int64, error) {

	id := generateUniqueId()

	newRequest := models.Request{
		Id:                 id,
		Type:               typeRequest,
		RequestingAdmin:    requestingAdminId,
		RequestedChapterId: requestedChapterId,
	}

	id, err := storage.RegisterRequest(db, newRequest)

	if err != nil {
		return 0, err
	}

	admins, err := storage.SelectAllSupeAdmins(db)

	for _, admin := range admins {
		ad, err := storage.SelectAdminWIthId(db, admin.Id)

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
