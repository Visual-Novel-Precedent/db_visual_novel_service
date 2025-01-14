package admin_

import (
	"db_novel_service/internal/services/admin"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type ChangeAdminStatusRequest struct {
	Id               int64           `json:"id"`
	Name             string          `json:"name"`
	Email            string          `json:"email"`
	Phone            string          `json:"phone"`
	Password         string          `json:"password"`
	AdminStatus      int             `json:"status"`
	CreatedChapters  []int64         `json:"create_chapters"`
	ChaptersProgress map[int64]int64 `json:"chapters_progress"`
}

func ChangeAdminHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req ChangeAdminStatusRequest
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

		// Здесь должна быть логика получения данных пользователя
		// Например, из базы данных:
		err = admin.ChangeAdmin(req.Id, req.Name, req.Email, req.Phone, req.Password, req.AdminStatus, req.CreatedChapters, req.ChaptersProgress, db)

		if err != nil {
			http.Error(w, "faik to change status", http.StatusInternalServerError)
		}

		// Формируем ответ
		response := map[string]interface{}{
			"id": req.Id,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
