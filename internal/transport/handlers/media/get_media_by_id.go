package media

import (
	"db_novel_service/internal/services/media"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type GetMediaByIdRequest struct {
	Id int64 `json:"id"`
}

func GetMediaByIdHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req DeleteMediaRequest
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

		media, err := media.GetMediaById(req.Id, db)

		if err != nil {
			http.Error(w, "fail to get media", http.StatusInternalServerError)
		}

		// Устанавливаем заголовки ответа
		w.Header().Set("Content-Type", media.ContentType)
		w.Header().Set("Content-Length", strconv.Itoa(len(media.FileData)))

		//// Формируем ответ
		//response := map[string]interface{}{
		//	"media": media,
		//}

		_, err = w.Write(media.FileData)
		if err != nil {
			log.Printf("Failed to write file: %v", err)
			return
		}
	}
}
