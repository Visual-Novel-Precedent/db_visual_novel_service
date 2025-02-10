package models

type Character struct {
	Id       int64 `gorm:"primary_key"`
	Name     string
	Slug     string
	Color    string
	Emotions map[int64]int64 `gorm:"type:json;column:emotions"` // индекс эмоции - id картинки
}
