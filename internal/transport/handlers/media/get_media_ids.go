package media

import (
	"db_novel_service/internal/services/media"
	"encoding/json"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"net/http"
)

func GetMediaHandler(db *gorm.DB, log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		media, err := media.GetMedia(db)

		if err != nil {
			http.Error(w, "fail to get media", http.StatusInternalServerError)
		}

		// Формируем ответ
		response := map[string]interface{}{
			"ids": media,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
