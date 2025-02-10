package node

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
)

func GetNodesById(
	id int64,
	db *gorm.DB,
) (models.Node, error) {

	node, err := storage.SelectNodeWIthId(db, id)

	return node, err
}
