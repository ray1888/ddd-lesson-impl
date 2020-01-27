package repository

import (
	"github.com/ray1888/ddd-lesson/v1/src/domain/student/model"
	"github.com/ray1888/ddd-lesson/v1/src/domain/student/model/read"
)

type StudentRepository interface {
	GetInfoById(id int) (read.StudentInfo, error)
	GetList() ([]read.Student, error)
	Add(studentInfo model.Student) (bool, error)
	Modify(studentInfo model.Student) (bool, error)
}
