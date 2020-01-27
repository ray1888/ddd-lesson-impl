package service

import (
	"github.com/ray1888/ddd-lesson/v1/src/domain/lesson/model"
	"github.com/ray1888/ddd-lesson/v1/src/domain/lesson/model/read"
	"github.com/ray1888/ddd-lesson/v1/src/domain/lesson/repository"
)

type ServiceInterface interface {
	GetLessonList() ([]*read.LessonReadModel, error)
	GetLesson(id int) (*read.LessonReadModel, error)
	GetLessonsByTeacherId(tid int) ([]*read.LessonReadModel, error)
	GetLessonsByStudentId(sid int) ([]*read.LessonReadModel, error)
	AddLesson(title string, teacherId, studentId int) (bool, error)
	DeleteLesson(id int) (bool, error)
	ModifyLesson(title string, id, studentId, teacherId int) (bool, error)
}

type Service struct {
	lesson repository.LessonRepository
}

func (s *Service) GetLessonList() ([]*read.LessonReadModel, error) {
	return s.lesson.GetList()
}

func (s *Service) GetLessonsByTeacherId(tid int) ([]*read.LessonReadModel, error) {
	if tid == 0 {
		return nil, nil
	}
	return s.lesson.GetByTeacherId(tid)
}

func (s *Service) GetLessonsByStudentId(sid int) ([]*read.LessonReadModel, error) {
	if sid == 0 {
		return nil, nil
	}
	return s.lesson.GetByStudentId(sid)
}

func (s *Service) GetLesson(id int) (*read.LessonReadModel, error) {
	if id == 0 {
		return nil, nil
	}
	return s.lesson.GetById(id)
}

func (s *Service) AddLesson(title string, studentId, teacherId int) (bool, error) {
	newLesson := new(model.Lesson)
	newLesson.Title = title
	newLesson.StudentId = studentId
	newLesson.TeacherId = teacherId
	return s.lesson.Add(newLesson)
}

func (s *Service) DeleteLesson(id int) (bool, error) {
	return s.lesson.Delete(id)
}

func (s *Service) ModifyLesson(title string, id, studentId, teacherId int) (bool, error) {
	newLesson := new(model.Lesson)
	newLesson.Title = title
	newLesson.Id = id
	newLesson.StudentId = studentId
	newLesson.TeacherId = teacherId
	return s.lesson.Modify(newLesson)
}
