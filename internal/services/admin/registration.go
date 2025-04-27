package admin

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/services/request"
	"db_novel_service/internal/storage"
	"errors"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

const (
	AdminNotFoundError = "admin data not found"

	DefaultAdminStatus = -1

	NoChapter                = 0
	RegisterAdminTypeRequest = 2
)

func Registration(email string, name string, password string, db *gorm.DB) (int64, error) {
	ad, err := storage.SelectAdminWIthEmail(db, email)

	if err != nil && ad.Id != 0 {
		return 0, errors.New("admin with this email is already exist")
	}

	//if err.Error() != AdminNotFoundError {
	//	log.Println(err, "ошибка получения админа")
	//	return 0, err
	//}

	id := generateUniqueId()

	newAdmin := models.Admin{
		Id:               id,
		Email:            email,
		Password:         password,
		Name:             name,
		AdminStatus:      DefaultAdminStatus,
		CreatedChapters:  []int64{},
		RequestSent:      []int64{},
		RequestsReceived: []int64{},
	}

	_, err = storage.RegisterAdmin(db, newAdmin)

	if err != nil {
		return 0, err
	}

	_, err = request.CreateRequest(id, RegisterAdminTypeRequest, NoChapter, db)

	log.Println(err)

	if err != nil {
		return 0, errors.New("fail to send registration requests to another admin")
	}

	_, err = storage.RegisterPlayer(db, models.Player{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Admin:    true,
	})

	if err != nil {
		return 0, errors.New("error to create player for admin")
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
