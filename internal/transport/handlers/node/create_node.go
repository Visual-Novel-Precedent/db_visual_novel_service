package node

import (
	"db_novel_service/internal/services/node"
	"encoding/json"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

type CreateNodeRequest struct {
	ChapterId string `json:"chapter_id"`
	Slug      string `json:"slug"`
}

func CreateNodeHandler(db *gorm.DB, log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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
		var req CreateNodeRequest
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

		idC, err := strconv.ParseInt(req.ChapterId, 10, 64)

		if err != nil {
			if err != nil {
				log.Println("ошибка конвертации")
				http.Error(w, "Failed to covert id", http.StatusInternalServerError)
				return
			}
		}

		id, err := node.CreateNode(idC, req.Slug, db)

		log.Println(err)

		if err != nil {
			http.Error(w, "fail to create node", http.StatusInternalServerError)
		}

		log.Println("newNce", utils.ToString(id))

		// Формируем ответ
		response := map[string]interface{}{
			"id": utils.ToString(id),
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}

type CreateNodeResponse struct {
	Id string
}
