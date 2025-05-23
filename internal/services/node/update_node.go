package node

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"log"
)

func UpdateNodeValue(
	id int64,
	slug string,
	events map[int]models.Event,
	music int64,
	background int64,
	branchingFlag bool,
	condition map[string]int64,
	endFlag bool,
	endResult string,
	endText string,
	comment string,
	db *gorm.DB,
) error {
	node, err := storage.SelectNodeWIthId(db, id)

	log.Println(events)

	if err != nil {
		return err
	}

	newNode := node

	if slug != "" {
		newNode.Slug = slug
	}

	if events != nil {
		log.Println(events, "eeevevevveveve")
		newNode.Events = events
		log.Println(newNode.Events, "hhhhhhhhh")
	}

	if music != 0 {
		newNode.Music = music
	}

	if background != 1 {
		newNode.Background = background
	}

	newNode.Branching.Flag = branchingFlag

	if condition != nil {
		newNode.Branching.Flag = true
		newNode.Branching.Condition = condition
	}

	newNode.End.Flag = endFlag

	if endResult != "" {
		newNode.End.EndResult = endResult
	}

	if endText != "" {
		newNode.End.EndText = endText
	}

	if comment != "" {
		newNode.Comment = comment
	}

	log.Println(newNode.Events)

	_, err = storage.UpdateNode(db, id, *newNode)

	return err
}
