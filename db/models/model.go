package models

type Model interface {
	FindById(int64)
	Create()
	Save()
}
