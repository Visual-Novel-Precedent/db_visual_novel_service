package admin

import (
	"db_novel_service/internal/services/admin"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type UserAuthorisationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AdminAuthorisationHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req UserAuthorisationRequest
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

		// Здесь должна быть логика получения данных пользователя
		// Например, из базы данных:
		user, err := admin.Authorization(req.Email, req.Password, db)

		// Формируем ответ
		response := map[string]interface{}{
			"id":               user.Id,
			"name":             user.Name,
			"adminStatus":      user.AdminStatus,
			"createdChapters":  user.CreatedChapters,
			"requestSent":      user.RequestSent,
			"requestsReceived": user.RequestsReceived,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
