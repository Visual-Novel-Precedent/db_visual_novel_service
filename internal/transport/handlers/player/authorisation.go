package player_

import (
	"db_novel_service/internal/services/player"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type PlayerAuthorisationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func PlayerAuthorisationHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req PlayerAuthorisationRequest
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
		player, err := player.Authorization(req.Email, req.Password, db)

		if err != nil {
			if err.Error() == "invalid password" {
				http.Error(w, "invalid password", http.StatusForbidden)
			}

			http.Error(w, "error to get user", http.StatusInternalServerError)
		}

		// Формируем ответ
		response := map[string]interface{}{
			"player": player,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
