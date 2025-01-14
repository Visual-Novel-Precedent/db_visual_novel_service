package admin_

import (
	"db_novel_service/internal/services/admin"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type AdminChapterProgressRequest struct {
	Id        int64 `json:"id"`
	ChapterId int64 `json:"chapter_id"`
	NodeId    int64 `json:"node_id"`
	EndFlag   bool  `json:"end_flag"`
}

func AdminChapterProgressHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req AdminChapterProgressRequest
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

		err = admin.UpdateChapterProgress(req.Id, req.ChapterId, req.NodeId, req.EndFlag, db)

		if err != nil {
			http.Error(w, "Fail to update chapter progress", http.StatusInternalServerError)
			return
		}
	}
}
