package request

import (
	"db_novel_service/internal/services/request"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type GetMyReceivedRequest struct {
	Id int64 `json:"id"`
}

func GetReceivedRequestHandler(db *gorm.DB) http.HandlerFunc {
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

		requests, err := request.GetReceivedRequests(req.Id, db)

		// Формируем ответ
		response := map[string]interface{}{
			"received_requests": requests,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
