package node

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"log"
)

func GetNodesByChapterId(
	id int64,
	db *gorm.DB,
) ([]models.Node, int64, error) {

	chapter, err := storage.SelectChapterWIthId(db, id)

	log.Println("nnnnnnnn", chapter.Nodes)

	var nodes []models.Node

	for _, n := range chapter.Nodes {
		log.Println("1")
		node, err := storage.SelectNodeWIthId(db, n)

		if err != nil {
			log.Println("no")
			return nil, 0, err
		}

		if node != nil {
			log.Println("yes")
			nodes = append(nodes, *node)
		}
	}

	log.Println(nodes, "nodes")

	return nodes, chapter.StartNode, err
}
