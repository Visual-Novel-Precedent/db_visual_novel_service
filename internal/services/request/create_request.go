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

	log.Println(requestingAdminId, typeRequest, requestedChapterId)

	newRequest := models.Request{
		Id:                 id,
		Type:               typeRequest,
		RequestingAdmin:    requestingAdminId,
		RequestedChapterId: requestedChapterId,
	}

	_, err := storage.RegisterRequest(db, newRequest)

	if err != nil {
		log.Println("ошибка регестрирования запроса")
		return 0, err
	}

	admin, err := storage.SelectAdminWithId(db, requestingAdminId)

	if err != nil {
		log.Println("ошибка обноружения админа")
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

	log.Println(typeRequest, "typeRequest")

	if typeRequest == 1 && requestedChapterId != 0 {

		chapter, err := storage.SelectChapterWIthId(db, requestedChapterId)

		if err != nil {
			return 0, err
		}

		chapter.Status = 2

		_, err = storage.UpdateChapter(db, chapter.Id, chapter)

		if err != nil {
			return 0, err
		}

		log.Println("yовый статус", chapter.Status)
	}

	return id, nil
}

func generateUniqueId() int64 {
	// Получаем текущее время в миллисекундах (48 бит)
	timestamp := time.Now().UnixMilli()

	// Генерируем 16 случайных бит
	random := rand.Int31n(1 << 16)

	// Объединяем timestamp и random в 64-битное число
	return (int64(timestamp) << 16) | int64(random)
}
