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

func DeleteRequest(db *gorm.DB, email string) (int64, error) {
	var deletedRequest models.Node
	result := db.Where("email = ?", email).Delete(&deletedRequest)
	if result.RowsAffected == 0 {
		return 0, errors.New("request data not update")
	}
	return result.RowsAffected, nil
}
