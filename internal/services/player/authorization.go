package player

import (
	"db_novel_service/internal/storage"
	"fmt"
	"gorm.io/gorm"
)

const (
	AdminStatus  = "admin"
	PlayerStatus = "player"
)

func Authorization(email string, password string, db *gorm.DB) (int64, map[int64]int64, string, string, error) {
	admin, err := storage.SelectAdminWIthEmail(db, email)

	if err == nil {
		if admin.Password != password {
			return -1, nil, "", "", fmt.Errorf("invalid password")
		}

		return admin.Id, admin.ChaptersProgress, admin.Name, AdminStatus, nil
	}

	user, err := storage.SelectPlayerWIthEmail(db, email)

	if err == nil {
		if user.Password != password {
			return -1, nil, "", "", fmt.Errorf("invalid password")
		}

		return user.Id, user.ChaptersProgress, PlayerStatus, user.Name, nil
	}

	return -1, nil, "", "", err
}
