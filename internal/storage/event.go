package storage

import (
	"db_novel_service/internal/models"
	"errors"
	"gorm.io/gorm"
)

func RegisterEvent(db *gorm.DB, event models.Event) (int64, error) {
	result := db.Create(&event)
	if result.RowsAffected == 0 {
		return 0, errors.New("character not created")
	}
	return result.RowsAffected, nil
}

func SelectEventWIthId(db *gorm.DB, id int64) (models.Event, error) {
	var character models.Event
	result := db.First(&character, "id = ?", id)
	if result.RowsAffected == 0 {
		return models.Event{}, errors.New("character data not found")
	}
	return character, nil
}

func UpdateEvent(db *gorm.DB, id int64, newEvent models.Event) (models.Event, error) {
	var event models.Event
	result := db.Model(&event).Where("id = ?", id).Updates(newEvent)
	if result.RowsAffected == 0 {
		return models.Event{}, errors.New("character data not update")
	}
	return event, nil
}

func DeleteEvent(db *gorm.DB, id int64) (int64, error) {
	var deletedEvent models.Event
	result := db.Where("id = ?", id).Delete(&deletedEvent)
	if result.RowsAffected == 0 {
		return 0, errors.New("event data not update")
	}
	return result.RowsAffected, nil
}
