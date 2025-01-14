package player_

import (
	player2 "db_novel_service/internal/services/player"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type PlayerRegistrationRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func PlayerRegistrationHandler(db *gorm.DB) http.HandlerFunc {
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

		id, err := player2.Registration(req.Email, req.Name, req.Password, db)

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
