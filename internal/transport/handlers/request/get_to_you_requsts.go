package request

import (
	"db_novel_service/internal/services/request"
	"encoding/json"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
)

type GetMyReceivedRequest struct {
	Id int64 `json:"id"`
}

func GetReceivedRequestHandler(db *gorm.DB, log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req GetMyRequestsRequest
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

		id, err := strconv.ParseInt(req.Id, 10, 64)

		if err != nil {
			if err != nil {
				log.Println("ошибка конвертации")
				http.Error(w, "Failed to covert id", http.StatusInternalServerError)
				return
			}
		}

		requests, err := request.GetReceivedRequests(id, db)

		// Формируем ответ
		response := map[string]interface{}{
			"received_requests": requests,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
