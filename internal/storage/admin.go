package storage

import (
	"database/sql"
	"db_novel_service/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	gorm "gorm.io/gorm"
)

func RegisterAdmin(db *gorm.DB, admin models.Admin) (int64, error) {
	result := db.Create(&admin)
	if result.RowsAffected == 0 {
		return 0, errors.New("admin not created")
	}
	return result.RowsAffected, nil
}

func SelectAdminWIthEmail(db *gorm.DB, email string) (models.Admin, error) {
	var admin models.Admin

	// Используем raw SQL с явной обработкой JSON
	query := `
        SELECT id, email, password,
               COALESCE(created_chapters::TEXT, '[]') as chapters_raw
        FROM admin
        WHERE email = $1
        LIMIT 1
    `

	row := db.Raw(query, email).Row()
	if err := row.Err(); err != nil {
		return models.Admin{}, err
	}

	var chaptersRaw sql.NullString
	err := row.Scan(&admin.Id, &admin.Email, &admin.Password, &chaptersRaw)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Admin{}, errors.New("admin data not found")
		}
		return models.Admin{}, err
	}

	// Десериализуем JSON в структуру
	var chaptersJSON []byte
	if chaptersRaw.Valid {
		chaptersJSON = []byte(chaptersRaw.String)
	} else {
		chaptersJSON = []byte("[]")
	}

	// Добавляем проверку на пустую строку
	if len(chaptersJSON) == 0 {
		chaptersJSON = []byte("[]")
	}

	err = json.Unmarshal(chaptersJSON, &admin.CreatedChapters)
	if err != nil {
		return models.Admin{}, fmt.Errorf("failed to unmarshal created_chapters: %w", err)
	}

	return admin, nil
}

//func SelectAdminWithId(db *gorm.DB, id int64) (models.Admin, error) {
//	var admin models.Admin
//	result := db.First(&admin, "id = ?", id)
//	if result.RowsAffected == 0 {
//		return models.Admin{}, errors.New("admin data not found")
//	}
//	return admin, nil
//}

func SelectAdminWithId(db *gorm.DB, id int64) (models.Admin, error) {
	var admin models.Admin

	// Используем raw SQL с явной обработкой JSON
	query := `
        SELECT id, email, password,
               COALESCE(created_chapters::TEXT, '[]') as chapters_raw
        FROM admin
        WHERE id = $1
        LIMIT 1
    `

	row := db.Raw(query, id).Row()
	if err := row.Err(); err != nil {
		return models.Admin{}, err
	}

	var chaptersRaw string
	err := row.Scan(&admin.Id, &admin.Email, &admin.Password, &chaptersRaw)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Admin{}, errors.New("admin data not found")
		}
		return models.Admin{}, err
	}

	// Десериализуем JSON в структуру
	err = json.Unmarshal([]byte(chaptersRaw), &admin.CreatedChapters)
	if err != nil {
		return models.Admin{}, fmt.Errorf("failed to unmarshal created_chapters: %w", err)
	}

	return admin, nil
}

func SelectAllSupeAdmins(db *gorm.DB) ([]models.Admin, error) {
	var admins []models.Admin
	result := db.Where("admin_status = ?", 1).Find(&admins)

	if result.RowsAffected == 0 {
		return nil, errors.New("no super admin found")
	}

	return admins, nil
}

func UpdateAdmin(db *gorm.DB, id int64, newAdmin models.Admin) (models.Admin, error) {
	var admin models.Admin

	chaptersJSON, err := json.Marshal(newAdmin.CreatedChapters)

	if err != nil {
		return models.Admin{}, err
	}

	requestsSendJSON, err := json.Marshal(newAdmin.RequestSent)

	if err != nil {
		return models.Admin{}, err
	}

	requestsReceivedJSON, err := json.Marshal(newAdmin.RequestsReceived)

	if err != nil {
		return models.Admin{}, err
	}

	result := db.Model(&admin).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"email":             newAdmin.Email,
			"password":          newAdmin.Password,
			"created_chapters":  json.RawMessage(chaptersJSON),
			"admin_status":      newAdmin.AdminStatus,
			"name":              newAdmin.Name,
			"request_sent":      json.RawMessage(requestsSendJSON),
			"requests_received": json.RawMessage(requestsReceivedJSON),
		})

	if result.RowsAffected == 0 {
		return models.Admin{}, errors.New("admin data not update")
	}

	return admin, nil
}
