package media

import (
	"db_novel_service/internal/services/media"
	"encoding/json"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

type CreateMediaRequest struct {
	File *multipart.FileHeader `json:"-"`
}

func CreateMediaHandler(db *gorm.DB, log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("получен запрос на добавление медиа")

		// Добавляем CORS заголовки
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")

		// Обработка предварительного запроса
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Проверка метода запроса
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Получаем Content-Type
		contentType := r.Header.Get("Content-Type")
		log.Printf("Content-Type: %s", contentType)

		var file io.Reader
		var err error

		// Проверяем тип загрузки
		if strings.HasPrefix(contentType, "multipart/form-data") {
			// Обработка multipart формата
			err = r.ParseMultipartForm(32 << 20)
			if err != nil {
				log.Printf("ошибка парсинга multipart: %v", err)
				http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
				return
			}

			// Получение файла из multipart формы
			fileHeader, _, err := r.FormFile("file")
			if err != nil {
				log.Printf("ошибка получения файла: %v", err)
				http.Error(w, "No file provided", http.StatusBadRequest)
				return
			}
			defer fileHeader.Close()

			file = fileHeader
		} else {
			// Прямая загрузка файла
			file = r.Body
		}

		// Проверка типа файла
		fileHeader := make([]byte, 512)
		_, err = file.Read(fileHeader)
		if err != nil {
			log.Printf("ошибка чтения заголовка файла: %v", err)
			http.Error(w, "Failed to read file header", http.StatusInternalServerError)
			return
		}

		// Возвращаем позицию файла на начало
		if closer, ok := file.(io.Seeker); ok {
			_, err = closer.Seek(0, io.SeekStart)
			if err != nil {
				log.Printf("ошибка возврата файла на начало: %v", err)
				http.Error(w, "Failed to seek file", http.StatusInternalServerError)
				return
			}
		}

		fileType := http.DetectContentType(fileHeader)

		if fileType != "image/png" && fileType != "audio/mpeg" {
			log.Printf("неподдерживаемый тип файла: %s", fileType)
			http.Error(w, "Only PNG images and MP3 audio files are allowed", http.StatusBadRequest)
			return
		}

		// Считываем весь файл
		fileData, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("ошибка чтения файла: %v", err)
			http.Error(w, "Failed to read file", http.StatusInternalServerError)
			return
		}

		// Сохраняем файл
		id, err := media.CreateMedia(fileData, "audio/mp3", db) // Изменено с fileType на "audio/mp3"
		if err != nil {
			log.Printf("ошибка сохранения файла: %v", err)
			http.Error(w, "Failed to save media", http.StatusInternalServerError)
			return
		}

		// Формируем ответ
		response := map[string]interface{}{
			"id": utils.ToString(id),
		}

		log.Println("id media", id)

		json.NewEncoder(w).Encode(response)
	}
}

type CreateMediaResponse struct {
	id string
}
