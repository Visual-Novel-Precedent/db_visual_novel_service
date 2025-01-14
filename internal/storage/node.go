package storage

import (
	"db_novel_service/internal/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func RegisterNode(db *gorm.DB, node models.Node) (int64, error) {
	result := db.Create(&node)
	if result.RowsAffected == 0 {
		return 0, errors.New("node not created")
	}
	return result.RowsAffected, nil
}

func SelectNodeWIthId(db *gorm.DB, id int64) (models.Node, error) {
	var node models.Node
	result := db.First(&node, "id = ?", id)
	if result.RowsAffected == 0 {
		return models.Node{}, errors.New("node data not found")
	}
	return node, nil
}

func UpdateNode(db *gorm.DB, id int64, newNode models.Node) (models.Node, error) {
	var node models.Node
	result := db.Model(&node).Where("id = ?", id).Updates(newNode)
	if result.RowsAffected == 0 {
		return models.Node{}, errors.New("node data not update")
	}
	return node, nil
}

func DeleteNode(db *gorm.DB, id int64) (int64, error) {
	var deletedNode models.Node
	result := db.Where("id = ?", id).Delete(&deletedNode)
	if result.RowsAffected == 0 {
		return 0, errors.New("event data not update")
	}
	return result.RowsAffected, nil
}

func GetNodeById(db *gorm.DB, id int64) (models.Node, error) {
	var node models.Node
	result := db.Where("id = ?", id).Find(&node)

	if result.Error != nil {
		return models.Node{}, fmt.Errorf("failed to fetch node: %v", result.Error)
	}

	return node, nil
}
