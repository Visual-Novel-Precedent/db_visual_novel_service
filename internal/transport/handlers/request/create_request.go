package request

import (
	"db_novel_service/internal/services/request"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type PlayerRegistrationRequest struct {
	RequestingAdminId string `json:"requesting_admin_id"`
	ChapterId         string `json:"chapter_id,omitempty"`
	Type              int    `json:"type"`
}

func CreateRequestHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("получили запрос на публикацию")
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
		var req PlayerRegistrationRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		log.Println(req)

		// Разбираем JSON
		err = json.Unmarshal(body, &req)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		reqId, err := strconv.ParseInt(req.RequestingAdminId, 10, 64)
		log.Println(req.RequestingAdminId, "req.RequestingAdminId")

		if err != nil {
			if err != nil {
				log.Println("ошибка конвертации id")
				http.Error(w, "Failed to covert id", http.StatusInternalServerError)
				return
			}
		}

		var chapterId int64

		if req.ChapterId != "" {
			chapterId, err = strconv.ParseInt(req.ChapterId, 10, 64)

			log.Println(req.ChapterId, "req.chapterId")

			if err != nil {
				if err != nil {
					log.Println("ошибка конвертации id")
					http.Error(w, "Failed to covert id", http.StatusInternalServerError)
					return
				}
			}
		}

		_, err = request.CreateRequest(reqId, req.Type, chapterId, db)

		if err != nil {
			http.Error(w, "fail to create request", http.StatusInternalServerError)
		}

		//// Формируем ответ
		//response := map[string]interface{}{
		//	"id": id,
		//}

		//// Отправляем ответ клиенту
		//json.NewEncoder(w).Encode(response)
	}
}
