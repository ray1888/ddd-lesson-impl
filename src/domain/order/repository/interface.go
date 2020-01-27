package repository

import (
	"github.com/ray1888/ddd-lesson/v1/src/domain/order/model"
	"github.com/ray1888/ddd-lesson/v1/src/domain/order/model/read"
)

type OrderRepository interface {
	GetInfo(id int) (*read.OrderRead, error)
	GetList() ([]*read.OrderRead, error)
	GetByStudentId(sid int) ([]*read.OrderRead, error)
	Add(item *model.Order) (bool, error)
	Delete(oid int) (bool, error)
	Modify(item *model.Order) (bool, error)
}
