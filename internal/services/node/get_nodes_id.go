package node

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/storage"
	"gorm.io/gorm"
	"log"
)

func GetNodesById(
	id int64,
	db *gorm.DB,
) (*models.Node, []models.Character, []models.Media, error) {

	node, err := storage.SelectNodeWIthId(db, id)

	log.Println(err)

	characters := make(map[int64]string)
	media := make(map[int64]string)

	media[node.Music] = ""

	for _, n := range node.Events {
		media[n.Sound] = ""
		for idCh, data := range n.CharactersInEvent {
			characters[idCh] = ""
			ch, _ := storage.SelectCharacterWIthId(db, idCh)
			for val, _ := range data {
				emotion := ch.Emotions[int64(val)]

				media[emotion] = ""
			}
		}
	}

	charactersRes := make([]models.Character, len(characters))
	mediaRes := make([]models.Media, len(media))

	log.Println(characters)

	for i, _ := range characters {

		log.Printf("character %s%", i)

		ch, err := storage.SelectCharacterWIthId(db, i)

		if err != nil {
			log.Printf(", %s", err)
		}

		if ch.Id != 0 {
			charactersRes = append(charactersRes, ch)
		}

		for j, k := range charactersRes {
			if k.Id == 0 {
				charactersRes = append(charactersRes[:j], charactersRes[j+1:]...)
			}
		}
	}

	for i, _ := range media {
		med, err := storage.SelectMediaWIthId(db, i)

		if err != nil {
			log.Println(err)
		}
		if med.Id != 0 {
			mediaRes = append(mediaRes, med)
		}

		for j, k := range mediaRes {
			if k.Id == 0 {
				mediaRes = append(mediaRes[:j], mediaRes[j+1:]...)
			}
		}
	}

	return node, charactersRes, mediaRes, err
}
