package models

type CourseModel struct {
	ID           int64
	Title        string
	Description  string
	Featured     bool
	CategoryId   int64
	CategoryName int64
}

func (a CourseModel) Validate() error {
	return nil
}
