package character

import (
	"db_novel_service/internal/services/character"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type CreateCharacterRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func CreateCharacterHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req CreateCharacterRequest
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

		id, err := character.CreateCharacter(req.Name, req.Slug, db)

		if err != nil {
			http.Error(w, "fail to create character", http.StatusInternalServerError)
		}

		// Формируем ответ
		response := map[string]interface{}{
			"id": id,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
