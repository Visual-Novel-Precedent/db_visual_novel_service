package node

import (
	"db_novel_service/internal/services/node"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

type GetNodeByIdRequest struct {
	Node int64 `json:"id"`
}

func GetNodeByIdHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req GetNodeByIdRequest
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

		log.Println(req.Node)

		node, characters, media, err := node.GetNodesById(req.Node, db)

		if err != nil {
			http.Error(w, "fail to get nodes", http.StatusInternalServerError)
		}

		// Формируем ответ
		response := map[string]interface{}{
			"node":       node,
			"characters": characters,
			"media":      media,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
