package storage

import (
	"db_novel_service/internal/models"
	"errors"
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
	result := db.First(&admin, "email = ?", email)
	if result.RowsAffected == 0 {
		return models.Admin{}, errors.New("admin data not found")
	}
	return admin, nil
}

func SelectAdminWIthId(db *gorm.DB, id int64) (models.Admin, error) {
	var admin models.Admin
	result := db.First(&admin, "id = ?", id)
	if result.RowsAffected == 0 {
		return models.Admin{}, errors.New("admin data not found")
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
	result := db.Model(&admin).Where("id = ?", id).Updates(newAdmin)
	if result.RowsAffected == 0 {
		return models.Admin{}, errors.New("admin data not update")
	}
	return admin, nil
}
