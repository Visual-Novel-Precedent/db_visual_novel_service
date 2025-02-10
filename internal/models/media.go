package models

type Media struct {
	Id   int64
	File []byte
	Type string // "music" - звук, "image" - картинки
}
