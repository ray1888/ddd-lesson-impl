package pg

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ray1888/ddd-lesson/v1/src/domain/lesson/model"
	"github.com/ray1888/ddd-lesson/v1/src/domain/lesson/model/read"
	"github.com/ray1888/ddd-lesson/v1/src/infrastructure/mapper"
)

type LessonRepositoryImpl struct {
	db *gorm.DB
}

func (lr *LessonRepositoryImpl) GetList() ([]*read.LessonReadModel, error) {
	row, err := lr.db.Table("lesson").Select("lesson.id,lesson.title, lesson.subject,lesson.grade_id" +
		"student.name, student.gender" +
		"teacher.name, teacher.gender").
		Joins("left join student on student.id=lesson.student_id").
		Joins("left join teacher on teacher.id=lesson.teacher_id").
		Rows()
	result := make([]*read.LessonReadModel, 0)
	if err != nil {
		return result, err
	}
	for row.Next() {
		item := new(read.LessonReadModel)
		err := row.Scan(item)
		if err != nil {
			continue
		}
		result = append(result, item)
	}
	return result, nil
}

func (lr *LessonRepositoryImpl) GetById(id int) ([]*read.LessonReadModel, error) {
	row, err := lr.db.Table("lesson").Select("lesson.id,lesson.title, lesson.subject,lesson.grade_id" +
		"student.name, student.gender" +
		"teacher.name, teacher.gender").
		Joins("left join student on student.id=lesson.student_id").
		Joins("left join teacher on teacher.id=lesson.teacher_id").
		Where(fmt.Sprintf("lesson.id = %d", id)).
		Rows()
	result := make([]*read.LessonReadModel, 0)
	if err != nil {
		return result, err
	}
	for row.Next() {
		item := new(read.LessonReadModel)
		err := row.Scan(item)
		if err != nil {
			continue
		}
		result = append(result, item)
	}
	return result, nil
}

func (lr *LessonRepositoryImpl) GetByTeacherId(tid int) ([]*read.LessonReadModel, error) {
	row, err := lr.db.Table("lesson").Select("lesson.id,lesson.title, lesson.subject,lesson.grade_id" +
		"student.name, student.gender" +
		"teacher.name, teacher.gender").
		Joins("left join student on student.id=lesson.student_id").
		Joins("left join teacher on teacher.id=lesson.teacher_id").
		Where(fmt.Sprintf("lesson.teacher_id = %d", tid)).
		Rows()
	result := make([]*read.LessonReadModel, 0)
	if err != nil {
		return result, err
	}
	for row.Next() {
		item := new(read.LessonReadModel)
		err := row.Scan(item)
		if err != nil {
			continue
		}
		result = append(result, item)
	}
	return result, nil
}

func (lr *LessonRepositoryImpl) GetByStudentId(sid int) ([]*read.LessonReadModel, error) {
	row, err := lr.db.Table("lesson").Select("lesson.id,lesson.title, lesson.subject,lesson.grade_id" +
		"student.name, student.gender" +
		"teacher.name, teacher.gender").
		Joins("left join student on student.id=lesson.student_id").
		Joins("left join teacher on teacher.id=lesson.teacher_id").
		Where(fmt.Sprintf("lesson.student_id = %d", sid)).
		Rows()
	result := make([]*read.LessonReadModel, 0)
	if err != nil {
		return result, err
	}
	for row.Next() {
		item := new(read.LessonReadModel)
		err := row.Scan(item)
		if err != nil {
			continue
		}
		result = append(result, item)
	}
	return result, nil
}

func (lr *LessonRepositoryImpl) Add(lesson *model.Lesson) (bool, error) {
	lr.db.Create(lesson)
	if lr.db.Error != nil {
		return false, lr.db.Error
	}
	return true, nil
}

func (lr *LessonRepositoryImpl) Update(lesson *model.Lesson) (bool, error) {
	lr.db.Update(lesson)
	if lr.db.Error != nil {
		return false, lr.db.Error
	}
	return true, nil
}

func (lr *LessonRepositoryImpl) Delete(id int) (bool, error) {
	lr.db.Delete((*mapper.LessonModel)(nil)).Where(fmt.Sprintf("lesson.id = %d", id))
	if lr.db.Error != nil {
		return false, lr.db.Error
	}
	return true, nil
}
