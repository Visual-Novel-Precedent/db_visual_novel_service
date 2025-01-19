package player_

import (
	"db_novel_service/internal/services/player"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type ChangePlayerStatusRequest struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	SoundSettings int    `json:"sound_settings"`
}

func ChangePlayerRequestHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req ChangePlayerStatusRequest
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
		err = player.ChangePlayer(req.Id, req.Name, req.Email, req.Phone, req.Password, req.SoundSettings, db)

		if err != nil {
			http.Error(w, "faik to change status", http.StatusInternalServerError)
		}
	}
}
