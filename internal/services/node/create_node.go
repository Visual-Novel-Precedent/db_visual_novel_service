package node

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

func CreateNode(chapterId int64, slug string, parentId int64, db *gorm.DB) (int64, error) {

	id := generateUniqueId()

	log.Println("новый узел")

	newNode := models.Node{
		Id:        id,
		Slug:      slug,
		ChapterId: chapterId,
		Events:    map[int]models.Event{},
		Branching: models.Branching{},
		End:       models.EndInfo{},
		Comment:   " ",
	}

	log.Println(newNode)

	chapter, err := storage.SelectChapterWIthId(db, chapterId)

	if err != nil {
		return 0, err
	}

	chapter.Nodes = append(chapter.Nodes, id)

	if err != nil {
		return 0, err
	}

	if parentId != -1 {
		log.Println("id родителя", parentId)
		parent, err := storage.SelectNodeWIthId(db, parentId)

		if err == nil && parent != nil {
			if parent.Branching.Flag && parent.Branching.Condition != nil {
				parent.Branching.Condition[slug] = id
			} else {
				parent.Branching.Flag = true
				parent.Branching.Condition = make(map[string]int64)
				parent.Branching.Condition[slug] = id
			}
		} else {
			log.Println("ошибка нахождения родителя", err, parent)
			return 0, err
		}

		_, err = storage.UpdateNode(db, parentId, *parent)
		if err != nil {
			log.Println("ошибка при обновлении node")
			return 0, err
		}
	}

	_, err = storage.UpdateChapter(db, chapterId, chapter)
	if err != nil {
		log.Println("ошибка при обновлнии главы")
		return 0, err
	}

	_, err = storage.RegisterNode(db, newNode)
	if err != nil {
		log.Println("ошибка при регистрации узла ")
		return 0, err
	}

	return id, nil
}

func generateUniqueId() int64 {
	// Получаем текущее время в миллисекундах (48 бит)
	timestamp := time.Now().UnixMilli()

	// Генерируем 16 случайных бит
	random := rand.Int31n(1 << 16)

	// Объединяем timestamp и random в 64-битное число
	return (int64(timestamp) << 16) | int64(random)
}
