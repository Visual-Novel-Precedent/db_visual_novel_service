package node

import (
	"db_novel_service/internal/models"
	nodeM "db_novel_service/internal/services/node"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type UpdateNodeRequest struct {
	Id         string         `json:"id"`
	Slug       string         `json:"slug,omitempty"`
	Events     map[int]EventR `json:"events,omitempty"`
	Music      string         `json:"music,omitempty"`
	Background string         `json:"background,omitempty"`
	Branching  Branching      `json:"branching,omitempty"`
	End        EndInfo        `json:"end,omitempty"`
}

type EventR struct {
	Type              string                       `json:"type"` // 0 - монолог героя или закадровый голос, 1- персонаж появилсяб 2 - перслнаж ушел, 3 - персонаж произносит речь
	Character         string                       `json:"character"`
	Sound             string                       `json:"sound"`
	CharactersInEvent map[string]map[string]string `json:"characters_in_event" gorm:"type:json"` // персонаж - эмоция + позиция относительно левого края
	Text              string                       `json:"text"`
}

type Branching struct {
	Flag      bool              `json:"branching_flag,omitempty"`
	Condition map[string]string `json:"condition, omitempty"`
}

type EndInfo struct {
	Flag      bool   `json:"end_flag,omitempty"`
	EndResult string `json:"end_result,omitempty"`
	EndText   string `json:"end_text,omitempty"`
}

func UpdateNodeHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("Получен запрос на изменение nodes")

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
			log.Println("в изменении nodes не тот метод", r.Method)
			http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		var req UpdateNodeRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("ошибка получния тела запроса nodes")
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		// Разбираем JSON
		err = json.Unmarshal(body, &req)
		if err != nil {
			log.Printf("ошибка десериализации nodes for update: %v", err)
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		log.Printf("events, %s", req.Events)

		id, err := strconv.ParseInt(req.Id, 10, 64)

		if err != nil {
			if err != nil {
				log.Println("ошибка конвертации")
				http.Error(w, "Failed to covert id", http.StatusInternalServerError)
				return
			}
		}

		music, err := strconv.ParseInt(req.Music, 10, 64)

		if err != nil {
			if err != nil {
				log.Println("ошибка конвертации")
				http.Error(w, "Failed to covert id", http.StatusInternalServerError)
				return
			}
		}

		background, err := strconv.ParseInt(req.Background, 10, 64)

		if err != nil {
			if err != nil {
				log.Println("ошибка конвертации")
				http.Error(w, "Failed to covert id", http.StatusInternalServerError)
				return
			}
		}

		condition := make(map[string]int64)

		for i, j := range req.Branching.Condition {
			c, err := strconv.ParseInt(j, 10, 64)

			if err != nil {
				if err != nil {
					log.Println("ошибка конвертации")
					http.Error(w, "Failed to covert id", http.StatusInternalServerError)
					return
				}
			}
			condition[i] = c
		}

		log.Println(id, req.Slug, req.Events, music, background, req.Branching.Flag, condition, req.End.Flag, req.End.EndResult, req.End.EndText)

		log.Println(req.Events, "до преобрахзования")

		events := make(map[int]models.Event)

		for k, h := range req.Events {

			log.Println("hh", h)

			typ, err := strconv.ParseInt(h.Type, 10, 64)

			if err != nil {
				if err != nil {
					log.Println("ошибка конвертации")
					http.Error(w, "Failed to covert id", http.StatusInternalServerError)
					return
				}
			}

			ch, err := strconv.ParseInt(h.Character, 10, 64)

			if err != nil {
				if err != nil {
					log.Println("ошибка конвертации")
					http.Error(w, "Failed to covert id", http.StatusInternalServerError)
					return
				}
			}

			sound, err := strconv.ParseInt(h.Sound, 10, 64)

			if err != nil {
				if err != nil {
					log.Println("ошибка конвертации")
					http.Error(w, "Failed to covert id", http.StatusInternalServerError)
					return
				}
			}

			charactersIn := make(map[int64]map[int64]int64)

			print("парс персонажей в сцене", h.CharactersInEvent)

			log.Println("event")
			log.Println(req.Events)

			for index, ev := range h.CharactersInEvent {
				emPos := make(map[int64]int64)

				for em, pos := range ev {

					emition, err := strconv.ParseInt(em, 10, 64)

					if err != nil {
						if err != nil {
							log.Println("ошибка конвертации")
							http.Error(w, "Failed to covert id", http.StatusInternalServerError)
							return
						}
					}

					position, err := strconv.ParseInt(pos, 10, 64)

					if err != nil {
						if err != nil {
							log.Println("ошибка конвертации")
							http.Error(w, "Failed to covert id", http.StatusInternalServerError)
							return
						}
					}

					emPos[emition] = position
					print("empos", emPos)
				}

				in, err := strconv.ParseInt(index, 10, 64)

				if err != nil {
					if err != nil {
						log.Println("ошибка конвертации")
						http.Error(w, "Failed to covert id", http.StatusInternalServerError)
						return
					}
				}

				charactersIn[in] = emPos
			}

			log.Println(h.Text)

			events[k] = models.Event{
				Type:              typ,
				Character:         ch,
				Sound:             sound,
				CharactersInEvent: charactersIn,
				Text:              h.Text,
			}
		}

		log.Println("events for update", events)

		err = nodeM.UpdateNodeValue(id, req.Slug, events, music, background, req.Branching.Flag, condition, req.End.Flag, req.End.EndResult, req.End.EndText, db)

		log.Println("node успешно обноален")

		if err != nil {
			log.Println("fail to update node")
			http.Error(w, "fail to update node", http.StatusInternalServerError)
		}
	}
}
