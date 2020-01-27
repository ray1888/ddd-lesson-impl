package repository

import (
	"github.com/ray1888/ddd-lesson/v1/src/domain/teacher/model"
	"github.com/ray1888/ddd-lesson/v1/src/domain/teacher/model/read"
)

type TeacherRepository interface {
	GetTeacherList() ([]read.Teacher, error)
	GetTeacherInfoById(id int) (read.Teacher, error)
	AddTeacher(teacherInfo model.Teacher) (bool, error)
}
