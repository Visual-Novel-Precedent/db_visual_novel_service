package models

type Request struct {
	Id                 int64
	Type               int // 0 - request for super admin, 1 - request for publication chapter, 2 - request for registration
	Status             int // 0 - on review, 1 - approved
	RequestingAdmin    int64
	RequestedChapterId int64
}
