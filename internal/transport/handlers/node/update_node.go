package node

import (
	nodeM "db_novel_service/internal/services/node"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type UpdateNodeRequest struct {
	Id         int64     `json:"id"`
	Slug       string    `json:"slug"`
	Events     []int64   `json:"events"`
	Music      string    `json:"music"`
	Background string    `json:"background"`
	Branching  Branching `json:"branching"`
	End        EndInfo   `json:"end"`
}

type Branching struct {
	Flag      bool          `json:"branching_flag"`
	Condition map[int]int64 `json:"condition"`
}

type EndInfo struct {
	Flag      bool   `json:"end_flag"`
	EndResult string `json:"end_result"`
	EndText   string `json:"end_text"`
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

		err = nodeM.UpdateNode(req.Id, req.Slug, req.Events, req.Music, req.Background, req.Branching.Flag, req.Branching.Condition, req.End.Flag, req.End.EndResult, req.End.EndText, db)

		if err != nil {
			http.Error(w, "fail to update node", http.StatusInternalServerError)
		}
	}
}
