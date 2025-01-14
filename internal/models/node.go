package models

type Node struct {
	Id         int64
	Slug       string
	Events     []int64
	ChapterId  int64
	Music      string
	Background string
	Branching  Branching
	End        EndInfo
}

type Branching struct {
	Flag      bool
	Condition map[int]int64 //Вариант и следующий узел
}

type EndInfo struct {
	Flag      bool
	EndResult string
	EndText   string
}
