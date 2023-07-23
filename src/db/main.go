package db

type Db[T any] interface {
	Add(item T) error
	GetById(id uint) (T, error)
	UpdateById(id uint, item T) error
}
