package admin_

import (
	"db_novel_service/internal/services/admin"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type AdminRegistrationRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func AdminRegistrationHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req AdminRegistrationRequest
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
		id, err := admin.Registration(req.Email, req.Name, req.Password, db)

		if err != nil {
			http.Error(w, "fail to register admin", http.StatusInternalServerError)
		}

		// Формируем ответ
		response := map[string]interface{}{
			"id": id,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
