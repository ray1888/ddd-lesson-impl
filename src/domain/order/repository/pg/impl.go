package pg

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/ray1888/ddd-lesson/v1/src/domain/order/model"
	"github.com/ray1888/ddd-lesson/v1/src/domain/order/model/read"
	"github.com/ray1888/ddd-lesson/v1/src/infrastructure/mapper"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func (or *OrderRepositoryImpl) GetList() ([]*read.OrderRead, error) {
	row, err := or.db.Select("order.id, order.price, order.student_id, order.paid_at from 'order' ").
		Rows()
	result := make([]*read.OrderRead, 0)
	if err != nil {
		return result, err
	}
	for row.Next() {
		item := new(read.OrderRead)
		err := row.Scan(item)
		if err != nil {
			continue
		}
		result = append(result, item)
	}
	return result, nil
}

func (or *OrderRepositoryImpl) GetInfo(oid int) ([]*read.OrderRead, error) {
	row, err := or.db.Select("order.id, order.price, order.student_id, order.paid_at from 'order' ").
		Where(fmt.Sprintf("order.id = %d", oid)).
		Rows()
	result := make([]*read.OrderRead, 0)
	if err != nil {
		return result, err
	}
	for row.Next() {
		item := new(read.OrderRead)
		err := row.Scan(item)
		if err != nil {
			continue
		}
		result = append(result, item)
	}
	return result, nil
}

func (or *OrderRepositoryImpl) GetByStudentId(sid int) ([]*read.OrderRead, error) {
	row, err := or.db.Select("order.id, order.price, order.student_id, order.paid_at from 'order' ").
		Where(fmt.Sprintf("order.student_id = %d", sid)).
		Rows()
	result := make([]*read.OrderRead, 0)
	if err != nil {
		return result, err
	}
	for row.Next() {
		item := new(read.OrderRead)
		err := row.Scan(item)
		if err != nil {
			continue
		}
		result = append(result, item)
	}
	return result, nil
}

func (or *OrderRepositoryImpl) Add(lesson *model.Order) (bool, error) {
	or.db.Create(lesson)
	if or.db.Error != nil {
		return false, or.db.Error
	}
	return true, nil
}

func (or *OrderRepositoryImpl) Modify(lesson *model.Order) (bool, error) {
	or.db.Update(lesson)
	if or.db.Error != nil {
		return false, or.db.Error
	}
	return true, nil
}

func (or *OrderRepositoryImpl) Delete(id int) (bool, error) {
	or.db.Delete((*mapper.OrderModel)(nil)).Where(fmt.Sprintf("order.id = %d", id))
	if or.db.Error != nil {
		return false, or.db.Error
	}
	return true, nil
}
