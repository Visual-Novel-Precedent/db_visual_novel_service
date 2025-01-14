package request

import (
	"db_novel_service/internal/services/request"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type RejectRequestRequest struct {
	IdRequest       int64 `json:"id_request"`
	IdApprovedAdmin int64 `json:"id_approved_admin"`
	IdReceivedAdmin int64 `json:"id_received_admin"`
}

func RejectRequestHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req GetMyRequestsRequest
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

		err = request.RejectRequest(req.Id, db)

		if err != nil {
			http.Error(w, "Failed to reject request", http.StatusInternalServerError)
		}
	}
}
