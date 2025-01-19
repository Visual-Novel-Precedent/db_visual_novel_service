package admin

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/services/request"
	"db_novel_service/internal/storage"
	"errors"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

const (
	AdminNotFoundError = "admin data not found"
	DefaultAdminStatus = -1

	NoChapter                = -1
	RegisterAdminTypeRequest = 1
)

func Registration(email string, name string, password string, db *gorm.DB) (int64, error) {
	_, err := storage.SelectAdminWIthEmail(db, email)

	if err == nil {
		return 0, errors.New("admin with this email is already exist")
	}

	if err.Error() != AdminNotFoundError {
		return 0, err
	}

	id := generateUniqueId()

	newAdmin := models.Admin{
		Id:          id,
		Email:       email,
		Password:    password,
		AdminStatus: DefaultAdminStatus,
	}

	_, err = storage.RegisterAdmin(db, newAdmin)

	if err != nil {
		return 0, err
	}

	_, err = request.CreateRequest(id, RegisterAdminTypeRequest, NoChapter, db)

	if err != nil {
		return 0, errors.New("fail to send registration requests to another admins")
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
	now := time.Now().UnixNano()
	random := rand.Int63n(1 << 32) // 32-битное случайное число
	return now ^ random
}
