package media

import (
	"db_novel_service/internal/services/media"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

type CreateMediaRequest struct {
	File *multipart.FileHeader `json:"-"`
}

func CreateMediaHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Проверяем Content-Type
		contentType := r.Header.Get("Content-Type")
		if contentType != "image/png" && contentType != "application/mp3" {
			http.Error(w, "Only PNG and MAP3 files are allowed", http.StatusBadRequest)
			return
		}

		// Читаем содержимое файла
		fileData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading file: %v", err)
			http.Error(w, "Failed to read file", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		id, err := media.CreateMedia(fileData, contentType, db)

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
