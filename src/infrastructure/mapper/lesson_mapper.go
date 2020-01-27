package mapper

import (
	"github.com/jinzhu/gorm"
)

type LessonModel struct {
	gorm.Model
	Title       string
	TeacherId   int
	TeacherName string
	StudentId   int
	StudentName string
	Subject     int
	GradeId     int
}
