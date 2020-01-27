package repository

import (
	"github.com/ray1888/ddd-lesson/v1/src/domain/teacher/model"
	"github.com/ray1888/ddd-lesson/v1/src/domain/teacher/model/read"
)

type TeacherRepository interface {
	GetList() ([]*read.Teacher, error)
	GetInfoById(id int) (*read.Teacher, error)
	Add(teacherInfo *model.Teacher) (bool, error)
	Modify(teacherInfo *model.Teacher) (bool, error)
}
