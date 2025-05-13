package player_

import (
	"db_novel_service/internal/services/player"
	"encoding/json"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"io/ioutil"
	"log"
	"net/http"
)

type PlayerAuthorisationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func PlayerAuthorisationHandler(db *gorm.DB, logger *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("Пришел запрос на авторизацию пользователя", r)

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
		var req PlayerAuthorisationRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("invalid body")
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		// Разбираем JSON
		err = json.Unmarshal(body, &req)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		log.Println("тело запроса на получеие игрока", req)

		// Здесь должна быть логика получения данных пользователя
		// Например, из базы данных:
		player, err := player.Authorization(req.Email, req.Password, db)

		if err != nil || player == nil {
			if err.Error() == "invalid password" {
				http.Error(w, "invalid password", http.StatusForbidden)
			}

			log.Println("player not found", err)
			http.Error(w, "error to get user", http.StatusInternalServerError)
		}

		if player == nil {
			http.Error(w, "error to get user", http.StatusForbidden)
		}

		// Формируем ответ
		response := map[string]interface{}{
			"Id":   utils.ToString(player.Id),
			"Name": player.Name,
		}

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}
