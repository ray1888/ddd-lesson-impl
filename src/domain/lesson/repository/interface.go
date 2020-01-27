package repository

import (
	. "github.com/ray1888/ddd-lesson/v1/src/domain/lesson/model/read"
)

type LessonRepository interface {
	GetLessonList() []LessonReadModel
	GetLessonById() LessonReadModel
}
