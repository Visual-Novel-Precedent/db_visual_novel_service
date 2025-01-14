package character

import (
	"db_novel_service/internal/services/character"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type UpdateCharacterRequest struct {
	Id       int64            `json:"id"`
	Name     string           `json:"name"`
	Slug     string           `json:"slug"`
	Color    string           `json:"color"`
	Emotions map[int64]string `json:"emotions"`
}

func UpdateCharacterHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req UpdateCharacterRequest
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

		err = character.UpdateCharacter(req.Id, req.Name, req.Slug, req.Color, req.Emotions, db)

		if err != nil {
			http.Error(w, "fail to create character", http.StatusInternalServerError)
		}
	}
}
