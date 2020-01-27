package mapper

import (
	"github.com/jinzhu/gorm"
)

type TeacherModel struct {
	gorm.Model
	Name       string `gorm:"type:varchar(100)"`
	Subject    int
	Gender     int
	CareerYear int
	Phone      string `gorm:"type:varchar(16)"`
}
