package node

import (
	"db_novel_service/internal/services/node"
	"db_novel_service/internal/transport/handlers/character"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type GetNodeByIdRequest struct {
	Node string `json:"id"`
}

func GetNodeByIdHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("Полуыен запрос на получение node")

		// Добавляем CORS заголовки
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")

		// Обрабатываем предварительный запрос (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

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

		id, err := strconv.ParseInt(req.Node, 10, 64)

		if err != nil {
			if err != nil {
				log.Println("ошибка конвертации")
				http.Error(w, "Failed to covert id", http.StatusInternalServerError)
				return
			}
		}

		log.Println(id)

		node, characters, media, err := node.GetNodesById(id, db)

		if err != nil {
			log.Println("Ошибка получения node", err)
			http.Error(w, "fail to get nodes", http.StatusInternalServerError)
		}

		var nodesRes ResponseNode

		if node != nil {
			nodesRes = convertToResponseNode(*node)
		}

		// Формируем ответ
		response := map[string]interface{}{
			"node":       nodesRes,
			"characters": character.PrepareCharacterForResponse(&characters),
			"media":      media,
		}

		log.Println("node успешно отправлен", nodesRes)

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
