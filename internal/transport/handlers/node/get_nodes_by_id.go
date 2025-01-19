package node

import (
	"db_novel_service/internal/services/node"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type GetNodeByChapterIdRequest struct {
	ChapterId int64 `json:"chapter_id"`
}

func GetNodeByChapterIdHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req GetNodeByChapterIdRequest
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

		nodes, err := node.GetNodesById(req.ChapterId, db)

		if err != nil {
			http.Error(w, "fail to get nodes", http.StatusInternalServerError)
		}

		// Формируем ответ
		response := map[string]interface{}{
			"nodes": nodes,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
