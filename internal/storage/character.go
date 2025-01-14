package storage

import (
	"db_novel_service/internal/models"
	"errors"
	"gorm.io/gorm"
)

func RegisterCharacter(db *gorm.DB, character models.Character) (int64, error) {
	result := db.Create(&character)
	if result.RowsAffected == 0 {
		return 0, errors.New("character not created")
	}
	return result.RowsAffected, nil
}

func SelectCharacterWIthId(db *gorm.DB, id int64) (models.Character, error) {
	var character models.Character
	result := db.First(&character, "id = ?", id)
	if result.RowsAffected == 0 {
		return models.Character{}, errors.New("character data not found")
	}
	return character, nil
}

func UpdateCharacter(db *gorm.DB, id int64, newCharacter models.Character) (models.Character, error) {
	var character models.Character
	result := db.Model(&character).Where("id = ?", id).Updates(newCharacter)
	if result.RowsAffected == 0 {
		return models.Character{}, errors.New("character data not update")
	}
	return character, nil
}

func DeleteCharacter(db *gorm.DB, id int64) (int64, error) {
	var deletedCharacter models.Chapter
	result := db.Where("id = ?", id).Delete(&deletedCharacter)
	if result.RowsAffected == 0 {
		return 0, errors.New("character data not update")
	}
	return result.RowsAffected, nil
}
