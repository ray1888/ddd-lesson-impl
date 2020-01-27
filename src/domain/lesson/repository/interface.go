package repository

import (
	"github.com/ray1888/ddd-lesson/v1/src/domain/lesson/model"
	. "github.com/ray1888/ddd-lesson/v1/src/domain/lesson/model/read"
)

type LessonRepository interface {
	GetList() ([]*LessonReadModel, error)
	GetById(id int) (*LessonReadModel, error)
	GetByTeacherId(tid int) ([]*LessonReadModel, error)
	GetByStudentId(sid int) ([]*LessonReadModel, error)
	Add(lesson *model.Lesson) (bool, error)
	Delete(id int) (bool, error)
	Modify(lesson *model.Lesson) (bool, error)
}
