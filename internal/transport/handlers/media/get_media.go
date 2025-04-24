package media

import (
	"db_novel_service/internal/services/media"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetMediaByIdGetHandler(db *gorm.DB, log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Добавляем CORS заголовки
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept")

		// Обрабатываем предварительный запрос (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Проверяем, что это GET-запрос
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Получаем ID из параметров URL
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			log.Printf("не указан параметр id")
			http.Error(w, "Missing required parameter: id", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			log.Printf("ошибка конвертации id: %v", err)
			http.Error(w, "Invalid id format", http.StatusBadRequest)
			return
		}

		media, err := media.GetMediaById(id, db)

		if err != nil {
			log.Printf("ошибка получения медиа: %v", err)
			http.Error(w, "fail to get media", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Length", strconv.Itoa(len(media.FileData)))
		w.Header().Set("Accept-Ranges", "bytes")

		_, err = w.Write(media.FileData)
		if err != nil {
			log.Printf("ошибка отправки данных: %v", err)
			return
		}

		log.Printf("медиа успешно отправлено (id: %d, тип: %s)", id, media.ContentType)
	}
}
