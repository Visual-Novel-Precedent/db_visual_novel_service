package media

import (
	"bytes"
	"db_novel_service/internal/services/media"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type GetMediaByIdRequest struct {
	Id string `json:"id"`
}

func GetMediaByIdHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Добавляем CORS заголовки
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")

		// Обрабатываем предварительный запрос (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req GetMediaByIdRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("ошибка чтения тела запроса: %v", err)
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		// Разбираем JSON
		err = json.Unmarshal(body, &req)
		if err != nil {
			log.Printf("ошибка разбора JSON: %v", err)
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(req.Id, 10, 64)
		if err != nil {
			log.Printf("ошибка конвертации id: %v", err)
			http.Error(w, "Failed to convert id", http.StatusBadRequest)
			return
		}

		media, err := media.GetMediaById(id, db)

		if err != nil {
			log.Printf("ошибка получения медиа: %v", err)
			http.Error(w, "fail to get media", http.StatusInternalServerError)
			return
		}

		var contentType string
		switch media.ContentType {
		case "mpga":
			contentType = "audio/mpeg"
		case "":
			// Если тип не указан, пытаемся определить по расширению
			if bytes.HasSuffix(media.FileData, []byte("ID3")) {
				contentType = "audio/mpeg"
			} else {
				contentType = "application/octet-stream"
			}
		default:
			contentType = media.ContentType
		}

		w.Header().Set("Content-Type", contentType)
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
