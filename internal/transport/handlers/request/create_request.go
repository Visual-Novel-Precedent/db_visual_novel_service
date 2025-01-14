package request

import (
	"db_novel_service/internal/services/request"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type PlayerRegistrationRequest struct {
	RequestingAdminId int64 `json:"requesting_admin_id"`
	ChapterId         int64 `json:"chapter_id"`
	Type              int   `json:"type"`
}

func CreateRequestHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		// Разбираем JSON
		err = json.Unmarshal(body, &req)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		id, err := request.CreateRequest(req.RequestingAdminId, req.Type, req.ChapterId, db)

		if err != nil {
			http.Error(w, "fail to create request", http.StatusInternalServerError)
		}

		// Формируем ответ
		response := map[string]interface{}{
			"id": id,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
