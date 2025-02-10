package chapter

import (
	"db_novel_service/internal/services/chapter"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type UpdateChapterRequest struct {
	Id             int64   `json:"id"`
	Name           string  `json:"name,omitempty"`
	StartNode      int64   `json:"start_node,omitempty"`
	Nodes          []int64 `json:"nodes,omitempty"`
	Characters     []int64 `json:"characters,omitempty"`
	Status         int     `json:"status,omitempty"` // 0 - черновик, 1 - на проверке, 2 - опубликована
	UpdateAuthorId int64   `json:"update_author_id,omitempty"`
}

func UpdateChapterHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req UpdateChapterRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		// Разбираем JSON
		err = json.Unmarshal(body, &req)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		err = chapter.UpdateChapter(req.Id, req.Name, req.Nodes, req.Characters, req.UpdateAuthorId, req.StartNode, req.Status, db)

		if err != nil {
			http.Error(w, "fail to create chapter", http.StatusInternalServerError)
		}
	}
}
