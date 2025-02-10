package models

type Event struct {
	Type              int // 0 - монолог героя или закадровый голос, 1- персонаж появилсяб 2 - перслнаж ушел, 3 - персонаж произносит речь
	Character         int64
	Sound             int64
	CharactersInEvent map[int64]map[int]int // персонаж - эмоция + позиция относительно левого края
	Text              string
}
