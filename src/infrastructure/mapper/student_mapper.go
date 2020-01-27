package mapper

import "github.com/jinzhu/gorm"

type StudentModel struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100)"`
	Subject int
	Gender  int
	GradeId int
	Phone   string `gorm:"type:varchar(16)"`
}
