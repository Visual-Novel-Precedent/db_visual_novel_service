package node

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
)

func GetNodesByChapterId(
	id int64,
	db *gorm.DB,
) ([]models.Node, int64, error) {

	chapter, err := storage.SelectChapterWIthId(db, id)

	var nodes []models.Node

	for _, n := range chapter.Nodes {
		node, err := storage.SelectNodeWIthId(db, n)

		if err != nil {
			return nil, 0, err
		}

		nodes = append(nodes, node)
	}

	return nodes, chapter.StartNode, err
}
