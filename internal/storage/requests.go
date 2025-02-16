package storage

import (
	"db_novel_service/internal/models"
	"errors"
	"gorm.io/gorm"
)

func RegisterRequest(db *gorm.DB, request models.Request) (int64, error) {
	result := db.Create(&request)
	if result.RowsAffected == 0 {
		return 0, errors.New("node not created")
	}
	return result.RowsAffected, nil
}

func SelectRequestWIthId(db *gorm.DB, id int64) (models.Request, error) {
	var request models.Request
	result := db.First(&request, "id = ?", id)
	if result.RowsAffected == 0 {
		return models.Request{}, errors.New("player data not found")
	}
	return request, nil
}

func UpdateRequest(db *gorm.DB, id int64, newRequest models.Request) (models.Request, error) {
	var request models.Request
	result := db.Model(&request).Where("id = ?", id).Updates(newRequest)
	if result.RowsAffected == 0 {
		return models.Request{}, errors.New("node data not update")
	}
	return request, nil
}

func DeleteRequest(db *gorm.DB, id int64) (int64, error) {
	var deletedRequest models.Request
	result := db.Where("id = ?", id).Delete(&deletedRequest)
	if result.RowsAffected == 0 {
		return 0, errors.New("request data not delete")
	}
	return result.RowsAffected, nil
}

func GetAllRequests(db *gorm.DB) ([]models.Request, error) {
	var requests []models.Request
	result := db.Find(&requests)

	if result.RowsAffected == 0 {
		return []models.Request{}, errors.New("записи не найдены")
	}

	return requests, nil
}
