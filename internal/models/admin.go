package models

type Admin struct {
	Id               int64 `gorm:"primary_key"`
	Name             string
	Email            string
	Password         string
	AdminStatus      int     // 0 - дефолтный алмин, 1 - сверхадмин, -1 - незаапрувенный админ
	CreatedChapters  []int64 `gorm:"type:json;column:created_chapters"`
	RequestSent      []int64 `gorm:"type:json;column:request_sent"`
	RequestsReceived []int64 `gorm:"type:json;column:requests_received"`
}
