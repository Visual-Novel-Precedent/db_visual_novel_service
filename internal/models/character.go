package models

type Character struct {
	ТС       int64
	Name     string
	Slug     string
	Color    string
	Emotions map[int64]string // индекс эмоции - url картинки
}
