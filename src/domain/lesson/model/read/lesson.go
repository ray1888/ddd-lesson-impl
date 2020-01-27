package read

type TeacherInfo struct {
	// Teacher Value Object
	TeacherName string
	Gender      int
}

type StudentInfo struct {
	// Student Value Object
	StudentName string
	Gender      int
}

type LessonListModel struct {
	Id      int
	Title   string
	Subject int
	GradeId int
	StudentInfo
	TeacherInfo
}
