package chapter

import (
	"db_novel_service/internal/services/node"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"time"
)

func UpdateChapter(
	id int64,
	name string,
	nodes []int64,
	characters []int64,
	updateAuthorId int64,
	startNode int64,
	status int,
	db *gorm.DB,
) error {
	chapter, err := storage.SelectChapterWIthId(db, id)

	if err != nil {
		return err
	}

	newChapter := chapter

	if newChapter.UpdatedAt == nil {
		newChapter.UpdatedAt = make(map[time.Time]int64)
	}

	newChapter.UpdatedAt[time.Now()] = updateAuthorId

	if name != "" {
		newChapter.Name = name
	}

	if nodes != nil {
		newChapter.Nodes = nodes

		deletedNodes := findDeletedElements(chapter.Nodes, nodes)

		if deletedNodes != nil && len(deletedNodes) > 0 {
			for _, nd := range deletedNodes {
				_, err = storage.DeleteNode(db, nd)

				if err != nil {
					return err
				}
			}

			currenNodes, _, err := node.GetNodesByChapterId(chapter.Id, db)

			if err != nil {
				return err
			}

			for _, cr := range currenNodes {
				if cr.Branching.Condition != nil {
					for cond, i := range cr.Branching.Condition {
						for _, del := range deletedNodes {
							if del == i {
								delete(cr.Branching.Condition, cond)

								_, _ = storage.UpdateNode(db, cr.Id, cr)
							}
						}
					}
				}
			}
		}
	}

	if characters != nil {
		newChapter.Characters = characters
	}

	if startNode != 0 {
		newChapter.StartNode = startNode
	}

	if status != 0 {
		newChapter.Status = status
	}

	_, err = storage.UpdateChapter(db, id, newChapter)

	return err
}

func findDeletedElements(original []int64, remaining []int64) []int64 {
	// Создаем мапу для хранения элементов оставшегося массива
	remainingMap := make(map[int64]bool)
	for _, num := range remaining {
		remainingMap[num] = true
	}

	// Находим удалённые элементы
	var deleted []int64
	for _, num := range original {
		if !remainingMap[num] {
			deleted = append(deleted, num)
		}
	}

	return deleted
}
