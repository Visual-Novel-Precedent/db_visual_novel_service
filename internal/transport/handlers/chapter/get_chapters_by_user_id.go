package chapter

import (
	"db_novel_service/internal/services/chapter"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type GetChaptersByUserIdRequest struct {
	UserId int64 `json:"user_id"`
}

func GetChaptersByUserIdHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req GetChaptersByUserIdRequest
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

		chapters, err := chapter.GetChaptersByUserId(db, req.UserId)

		if err != nil {
			http.Error(w, "fail to get chapters", http.StatusInternalServerError)
		}

		// Формируем ответ
		response := map[string]interface{}{
			"chapters": chapters,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
