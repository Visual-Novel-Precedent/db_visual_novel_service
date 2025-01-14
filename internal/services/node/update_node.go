package node

import (
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
)

func UpdateNode(
	id int64,
	slug string,
	events []int64,
	music string,
	background string,
	branchingFlag bool,
	condition map[int]int64,
	endFlag bool,
	endResult string,
	endText string,
	db *gorm.DB,
) error {
	node, err := storage.SelectNodeWIthId(db, id)

	if err != nil {
		return err
	}

	newNode := node

	if slug != "" {
		newNode.Slug = slug
	}

	if events != nil {
		newNode.Events = events
	}

	if music != "" {
		newNode.Music = music
	}

	if background != "" {
		newNode.Background = background
	}

	newNode.Branching.Flag = branchingFlag

	if condition != nil {
		newNode.Branching.Condition = condition
	}

	newNode.End.Flag = endFlag

	if endResult != "" {
		newNode.End.EndResult = endResult
	}

	if endText != "" {
		newNode.End.EndText = endText
	}

	_, err = storage.UpdateNode(db, id, newNode)

	return err
}
