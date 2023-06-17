package entities

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);uniqueIndex"`
	Description string `gorm:"type:text"`
	Level       string `gorm:"type:varchar(100)"`
	Point       int
	TestCode    string `gorm:"type:text"`
	ProblemCode string `gorm:"type:text"`
	Order       int16
}
