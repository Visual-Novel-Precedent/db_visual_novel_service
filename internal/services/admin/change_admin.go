package admin

import (
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
)

const (
	NilStatus = -1
)

func ChangeAdmin(
	id int64,
	name string,
	email string,
	phone string,
	password string,
	adminStatus int,
	createdChapters []int64,
	chaptersProgress map[int64]int64,
	db *gorm.DB,
) error {
	user, err := storage.SelectAdminWIthId(db, id)

	if err != nil {
		return err
	}

	if name != "" {
		user.Name = name
	}

	if email != "" {
		user.Email = email
	}

	if phone != "" {
		user.Phone = phone
	}

	if password != "" {
		user.Password = password
	}

	if adminStatus != NilStatus {
		user.AdminStatus = adminStatus
	}

	if createdChapters != nil {
		user.CreatedChapters = append(user.CreatedChapters, createdChapters...)
	}

	if chaptersProgress != nil {
		user.ChaptersProgress = chaptersProgress
	}

	_, err = storage.UpdateAdmin(db, id, user)

	return err
}
