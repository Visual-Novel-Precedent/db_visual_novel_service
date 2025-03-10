package node

import (
	"db_novel_service/internal/models"
	"db_novel_service/internal/services/node"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type GetNodeByChapterIdRequest struct {
	ChapterId string `json:"chapter_id"`
}

func GetNodeByChapterIdHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("Получен запрос на получение nodes")

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
		var req GetNodeByChapterIdRequest
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

		id, err := strconv.ParseInt(req.ChapterId, 10, 64)

		if err != nil {
			if err != nil {
				log.Println("ошибка конвертации")
				http.Error(w, "Failed to covert id", http.StatusInternalServerError)
				return
			}
		}

		node, startNode, err := node.GetNodesByChapterId(id, db)

		if err != nil {
			http.Error(w, "fail to get nodes", http.StatusInternalServerError)
		}

		log.Println("Ошибка получения nodes", err)

		// Формируем ответ
		response := map[string]interface{}{
			"nodes":      prepareResponseNodeList(node),
			"start_node": startNode,
		}

		log.Println("nodes успешно отправлен", len(prepareResponseNodeList(node)))

		// Отправляем ответ клиенту
		json.NewEncoder(w).Encode(response)
	}
}

type ResponseNode struct {
	Id         string `gorm:"primary_key"`
	Slug       string
	Events     map[int]ResponseEvent `gorm:"type:json"`
	ChapterId  string
	Music      string
	Background string
	Branching  ResponseBranching `gorm:"type:json"`
	End        ResponseEndInfo   `gorm:"type:json"`
	Comment    string
}

type ResponseEvent struct {
	Type              string                       `json:"type"` // 0 - монолог героя или закадровый голос, 1- персонаж появилсяб 2 - перслнаж ушел, 3 - персонаж произносит речь
	Character         string                       `json:"character"`
	Sound             string                       `json:"sound"`
	CharactersInEvent map[string]map[string]string `json:"characters_in_event"` // персонаж - эмоция + позиция относительно левого края
	Text              string                       `json:"text"`
}

type ResponseBranching struct {
	Flag      bool              `json:"branching_flag,omitempty"`
	Condition map[string]string `json:"condition, omitempty"`
}

type ResponseEndInfo struct {
	Flag      bool   `json:"end_flag,omitempty"`
	EndResult string `json:"end_result,omitempty"`
	EndText   string `json:"end_text,omitempty"`
}

func prepareResponseNodeList(nodes []models.Node) []ResponseNode {
	var res []ResponseNode

	for _, node := range nodes {
		res = append(res, convertToResponseNode(node))
	}

	return res
}

func convertToResponseNode(node models.Node) ResponseNode {
	responseEvents := make(map[int]ResponseEvent)

	// Конвертируем события
	for id, event := range node.Events {
		responseEvents[id] = ResponseEvent{
			Type:              ToString(event.Type),
			Character:         ToString(event.Character),
			Sound:             ToString(event.Sound),
			Text:              event.Text,
			CharactersInEvent: convertCharactersMap(event.CharactersInEvent),
		}
	}

	return ResponseNode{
		Id:         ToString(node.Id),
		Slug:       node.Slug,
		Events:     responseEvents,
		ChapterId:  ToString(node.ChapterId),
		Music:      ToString(node.Music),
		Background: ToString(node.Background),
		Branching: ResponseBranching{
			Flag:      node.Branching.Flag,
			Condition: convertConditionMap(node.Branching.Condition),
		},
		End: ResponseEndInfo{
			Flag:      node.End.Flag,
			EndResult: node.End.EndResult,
			EndText:   node.End.EndText,
		},
		Comment: node.Comment,
	}
}

// Вспомогательные функции для конвертации карт
func convertCharactersMap(source map[int64]map[int64]int64) map[string]map[string]string {
	result := make(map[string]map[string]string)
	for k, innerMap := range source {
		strInnerMap := make(map[string]string)
		for innerK, v := range innerMap {
			strInnerMap[ToString(innerK)] = ToString(v)
		}
		result[ToString(k)] = strInnerMap
	}
	return result
}

func convertConditionMap(source map[string]int64) map[string]string {
	result := make(map[string]string)
	for k, v := range source {
		result[k] = ToString(v)
	}
	return result
}

// Утилитарная функция для конвертации int64 в string
func ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
