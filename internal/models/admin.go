package models

type Admin struct {
	Id               int64
	Name             string
	Email            string
	Phone            string
	Password         string
	AdminStatus      int             // 0 - дефолтный алмин, 1 - сверхадмин, -1 - незаапрувенный админ
	CreatedChapters  []int64         // пройденные главы
	ChaptersProgress map[int64]int64 // Мапа id главы - id узла
	RequestSent      []int64
	RequestsReceived []int64
}
