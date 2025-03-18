package request

import (
	"db_novel_service/internal/services/request"
	"encoding/json"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ApproveRequestRequest struct {
	IdRequest string `json:"id_request"`
}

func ApproveRequestHandler(db *gorm.DB, log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("получен запрос на аппрув запроса")
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
		var req ApproveRequestRequest
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

		id, err := strconv.ParseInt(req.IdRequest, 10, 64)

		if err != nil {
			if err != nil {
				log.Println("ошибка конвертации id")
				http.Error(w, "Failed to covert id", http.StatusInternalServerError)
				return
			}
		}

		err = request.ApproveRequest(id, db)

		log.Println(err)

		if err != nil {
			http.Error(w, "Failed to approve request", http.StatusInternalServerError)
		}
	}
}
