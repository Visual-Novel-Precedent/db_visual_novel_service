package storage

import (
	"db_novel_service/internal/models"
	"errors"
	"gorm.io/gorm"
)

func RegisterPlayer(db *gorm.DB, player models.Player) (int64, error) {
	result := db.Create(&player)
	if result.RowsAffected == 0 {
		return 0, errors.New("node not created")
	}
	return result.RowsAffected, nil
}

func SelectPlayerWIthEmail(db *gorm.DB, email string) (models.Player, error) {
	var player models.Player
	result := db.First(&player, "email = ?", email)
	if result.RowsAffected == 0 {
		return models.Player{}, errors.New("player data not found")
	}
	return player, nil
}

func SelectPlayerWIthId(db *gorm.DB, id int64) (models.Player, error) {
	var player models.Player
	result := db.First(&player, "id = ?", id)
	if result.RowsAffected == 0 {
		return models.Player{}, errors.New("player data not found")
	}
	return player, nil
}

func UpdatePlayer(db *gorm.DB, id int64, newPlayer models.Player) (models.Player, error) {
	var player models.Player
	result := db.Model(&player).Where("id = ?", id).Updates(newPlayer)
	if result.RowsAffected == 0 {
		return models.Player{}, errors.New("node data not update")
	}
	return player, nil
}

func DeletePlayer(db *gorm.DB, email string) (int64, error) {
	var deletedPlayer models.Node
	result := db.Where("email = ?", email).Delete(&deletedPlayer)
	if result.RowsAffected == 0 {
		return 0, errors.New("player data not update")
	}
	return result.RowsAffected, nil
}
