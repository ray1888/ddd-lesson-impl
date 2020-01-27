package repository

import (
	"github.com/ray1888/ddd-lesson/v1/src/domain/student/model"
	"github.com/ray1888/ddd-lesson/v1/src/domain/student/model/read"
)

type StudentRepository interface {
	GetStudentInfoById(id int) (read.StudentInfo, error)
	GetStudentList() ([]read.Student, error)
	AddStudent(studentInfo model.Student) (bool, error)
	ModifyStudentInfo(studentInfo model.Student) (bool, error)
}
