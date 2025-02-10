package node

import (
	"db_novel_service/internal/models"
	nodeM "db_novel_service/internal/services/node"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type UpdateNodeRequest struct {
	Id         int64                `json:"id"`
	Slug       string               `json:"slug,omitempty"`
	Events     map[int]models.Event `json:"events,omitempty"`
	Music      int64                `json:"music,omitempty"`
	Background int64                `json:"background,omitempty"`
	Branching  Branching            `json:"branching,omitempty"`
	End        EndInfo              `json:"end,omitempty"`
}

type Branching struct {
	Flag      bool             `json:"branching_flag,omitempty"`
	Condition map[string]int64 `json:"condition, omitempty"`
}

type EndInfo struct {
	Flag      bool   `json:"end_flag,omitempty"`
	EndResult string `json:"end_result,omitempty"`
	EndText   string `json:"end_text,omitempty"`
}

func UpdateNodeHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req UpdateNodeRequest
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

		err = nodeM.UpdateNodeValue(req.Id, req.Slug, req.Events, req.Music, req.Background, req.Branching.Flag, req.Branching.Condition, req.End.Flag, req.End.EndResult, req.End.EndText, db)

		if err != nil {
			http.Error(w, "fail to update node", http.StatusInternalServerError)
		}
	}
}
