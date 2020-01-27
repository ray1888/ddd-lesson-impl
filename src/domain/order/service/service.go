package service

import (
	"time"

	lrepository "github.com/ray1888/ddd-lesson/v1/src/domain/lesson/repository"
	"github.com/ray1888/ddd-lesson/v1/src/domain/order/model"
	"github.com/ray1888/ddd-lesson/v1/src/domain/order/model/read"
	"github.com/ray1888/ddd-lesson/v1/src/domain/order/repository"
)

type ServiceInterface interface {
	GetOrderList() ([]*read.OrderRead, error)
	GetOrderInfo(oid int) (*read.OrderRead, error)
	GetOrderListByStudentId(sid int) ([]*read.OrderRead, error)
	AddOrder(sid int, price float32) (bool, error)
	ModifyOrder(oid, sid int, price float32) (bool, error)
	DeleteOrder(oid int) (bool, error)
}

const (
	DisconutNumber = 0.8
)

type Service struct {
	// Add Order needed to get lesson num to make discount
	OrderRepo  repository.OrderRepository
	LessonRepo lrepository.LessonRepository
}

func (s *Service) GetOrderList() ([]*read.OrderRead, error) {
	return s.OrderRepo.GetList()
}

func (s *Service) GetOrderInfo(oid int) (*read.OrderRead, error) {
	if oid == 0 {
		return nil, nil
	}
	return s.OrderRepo.GetInfo(oid)
}

func (s *Service) GetOrderListByStudentId(sid int) ([]*read.OrderRead, error) {
	if sid == 0 {
		return nil, nil
	}
	return s.OrderRepo.GetByStudentId(sid)
}

func (s *Service) AddOrder(sid int, price float32) (bool, error) {
	// Need to get lesson to make disconunt
	newOrder := new(model.Order)
	newOrder.StudentId = sid
	newOrder.PaidAt = time.Now()
	if s.getDisconuntByCountLesson(sid) != 0 {
		newOrder.Price = price * DisconutNumber
	} else {
		newOrder.Price = price
	}
	return s.OrderRepo.Add(newOrder)
}

func (s *Service) getDisconuntByCountLesson(sid int) int {
	lesson, err := s.LessonRepo.GetByStudentId(sid)
	if err != nil {
		return 0
	}
	return len(lesson)
}
