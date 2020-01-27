package service

import "github.com/ray1888/ddd-lesson/v1/src/domain/lesson/model/read"

type Service interface {
	GetStudentInfoById(id int) (read.StudentInfo, error)
	GetStudentList() ([]read.StudentInfo, error)
	AddStudent(StudentName, Phone string, Grade int) (bool, error)
	ModifyStudent(StudentName, Phone string, Id, Grade int) (bool, error)
}
