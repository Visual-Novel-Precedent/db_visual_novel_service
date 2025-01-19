package character

import (
	"db_novel_service/internal/services/character"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

func GetCharacterHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		characters, err := character.GetCharacters(db)

		if err != nil {
			http.Error(w, "fail to get characters", http.StatusInternalServerError)
		}

		// Формируем ответ
		response := map[string]interface{}{
			"characters": characters,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
