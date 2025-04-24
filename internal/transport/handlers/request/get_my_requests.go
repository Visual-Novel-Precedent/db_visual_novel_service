package request

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/services/request"
	"encoding/json"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

type GetMyRequestsRequest struct {
	Id string `json:"id"`
}

func GetMyRequestHandler(db *gorm.DB, log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("запрос на получение запросов получен")
		// Добавляем CORS заголовки
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")

		// Обрабатываем предварительный запрос (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		log.Println("корсы в запросах прошли ")

		// Проверяем, что это POST-запрос
		if r.Method != http.MethodPost {
			log.Println("некорректный метод")
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req GetMyRequestsRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("ошибка получния тела запроса в запросах")
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		// Разбираем JSON
		err = json.Unmarshal(body, &req)
		if err != nil {
			log.Println("ошибка в форматировании Json в запросах")
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(req.Id, 10, 64)

		if err != nil {
			if err != nil {
				log.Println("ошибка конвертации id")
				http.Error(w, "Failed to covert id", http.StatusInternalServerError)
				return
			}
		}

		requests, err := request.GetMyRequests(id, db)
		log.Println("проблема с получением запросов", err)

		print("запросы", requests)

		// Формируем ответ
		response := map[string]interface{}{
			"my_requests": prepareRequestResponse(requests),
		}

		log.Println("запросы после конвертации", prepareRequestResponse(requests))

		log.Println("запросы успешно отправлены")

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}

type ResponseRequest struct {
	Id                     string
	Type                   int
	Status                 int
	RequestingAdmin        string
	RequestedChapterId     string
	RequestingAdminName    string
	RequestedChapterIdName string
}

func prepareRequestResponse(requests []models.Request) []ResponseRequest {
	var res []ResponseRequest

	for _, r := range requests {
		res = append(res, ResponseRequest{
			Id:                     utils.ToString(r.Id),
			Type:                   r.Type,
			Status:                 r.Status,
			RequestingAdmin:        utils.ToString(r.RequestingAdmin),
			RequestedChapterId:     utils.ToString(r.RequestedChapterId),
			RequestingAdminName:    r.RequestedAdminName,
			RequestedChapterIdName: r.RequestedChapterName,
		})
	}

	return res
}
