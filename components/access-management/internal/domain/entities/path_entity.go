package entities

type PathEntity struct {
	ID      int64
	Title   string
	Courses []CourseEntity
}
