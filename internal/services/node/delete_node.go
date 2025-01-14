package node

import (
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
)

func DeleteNode(
	id int64,
	db *gorm.DB,
) error {
	node, err := storage.SelectNodeWIthId(db, id)

	if err != nil {
		return err
	}

	_, err = storage.DeleteNode(db, node.Id)

	return err
}
