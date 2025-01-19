package node

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
)

func GetNodesById(
	id int64,
	db *gorm.DB,
) ([]models.Node, error) {
	var nodes []models.Node

	node, err := storage.GetNodeById(db, id)

	if err != nil {
		return nil, err
	}

	nodes = append(nodes, node)

	return nodes, nil
}
