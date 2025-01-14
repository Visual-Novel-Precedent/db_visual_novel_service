package storage

import (
	"db_novel_service/internal/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func RegisterChapter(db *gorm.DB, chapter models.Chapter) (int64, error) {
	result := db.Create(&chapter)
	if result.RowsAffected == 0 {
		return 0, errors.New("chapter not created")
	}
	return result.RowsAffected, nil
}

func SelectChapterWIthId(db *gorm.DB, id int64) (models.Chapter, error) {
	var chapter models.Chapter
	result := db.First(&chapter, "id = ?", id)
	if result.RowsAffected == 0 {
		return models.Chapter{}, errors.New("chapter data not found")
	}
	return chapter, nil
}

func UpdateChapter(db *gorm.DB, id int64, newChapter models.Chapter) (models.Chapter, error) {
	var chapter models.Chapter
	result := db.Model(&chapter).Where("id = ?", id).Updates(newChapter)
	if result.RowsAffected == 0 {
		return models.Chapter{}, errors.New("chapter data not update")
	}
	return chapter, nil
}

func DeleteChapter(db *gorm.DB, id int64) (int64, error) {
	var deletedChapter models.Chapter
	result := db.Where("id = ?", id).Delete(&deletedChapter)
	if result.RowsAffected == 0 {
		return 0, errors.New("payment data not update")
	}
	return result.RowsAffected, nil
}

func GetChaptersForAdmin(db *gorm.DB) ([]models.Chapter, error) {
	var chapters []models.Chapter
	result := db.Find(&chapters)

	if result.Error != nil {
		return nil, result.Error
	}

	return chapters, nil
}

func FindPublishedChapters(db *gorm.DB) ([]models.Chapter, error) {
	var chapters []models.Chapter
	result := db.Where("status = ?", 2).Find(&chapters)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch published chapters: %v", result.Error)
	}

	if len(chapters) == 0 {
		return nil, fmt.Errorf("no published chapters found")
	}

	return chapters, nil
}
