package models

import (
	"time"
)

type Chapter struct {
	Id         int64
	Name       string
	Nodes      []int64
	Characters []int64
	Status     int // 0 - черновик, 1 - на проверке, 2 - опубликована
	UpdatedAt  map[time.Time]int64
	Author     int64
}
