package player

import (
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
)

func ChangePlayer(
	id int64,
	name string,
	email string,
	phone string,
	password string,
	db *gorm.DB,
) error {
	user, err := storage.SelectPlayerWIthId(db, id)

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

	_, err = storage.UpdatePlayer(db, id, user)

	return err
}
