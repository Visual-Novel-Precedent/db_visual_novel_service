package player_

import (
	"db_novel_service/internal/services/player"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

type ChangePlayerRequest struct {
	Id            int64  `json:"id"`
	Name          string `json:"name,omitempty"`
	Email         string `json:"email,omitempty"`
	Password      string `json:"password,omitempty"`
	SoundSettings int    `json:"sound_settings,omitempty"`
}

func ChangePlayerRequestHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req ChangePlayerRequest
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
		err = player.ChangePlayer(req.Id, req.Name, req.Email, req.Password, req.SoundSettings, db)

		log.Println(err)

		if err != nil {
			http.Error(w, "fail to change status", http.StatusInternalServerError)
		}
	}
}
