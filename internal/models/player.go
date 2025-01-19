package models

type Player struct {
	Id                int64
	Name              string
	Email             string
	Phone             string
	Password          string
	Admin             bool
	CompletedChapters []int64         // пройденные главы
	ChaptersProgress  map[int64]int64 // Мапа id главы - id узла
	SoundSettings     int             //percent
}
