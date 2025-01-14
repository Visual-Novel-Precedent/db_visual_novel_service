package models

type Character struct {
	Id       int64
	Name     string
	Slug     string
	Color    string
	Emotions map[int64]string // индекс эмоции - url картинки
}
