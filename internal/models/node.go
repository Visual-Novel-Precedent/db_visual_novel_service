package models

type Node struct {
	Id         int64 `gorm:"primary_key"`
	Slug       string
	Events     map[int]Event
	ChapterId  int64
	Music      int64
	Background int64
	Branching  Branching
	End        EndInfo
}

type Branching struct {
	Flag      bool
	Condition map[string]int64 //Вариант и следующий узел
}

type EndInfo struct {
	Flag      bool
	EndResult string
	EndText   string
}
