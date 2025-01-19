package models

type Event struct {
	Index              int
	Type               int // 0 - монолог героя или закадровый голос, 1- персонаж появилсяб 2 - перслнаж ушел, 3 - персонаж произносит речь
	CharacterArrived   Character
	CharacterDeparture Character
	Sound              string
	CharactersInEvent  map[int64]map[float64]float64 // позиция - персонаж
	Text               string
}
