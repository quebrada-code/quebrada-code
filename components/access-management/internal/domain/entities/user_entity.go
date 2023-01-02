package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID               uint `gorm:"primaryKey"`
	Name             string
	Email            string `gorm:"uniqueIndex"`
	EmailConfirmed   bool
	VerificationCode string
	ResetToken       string
	PasswordHash     string
	Password         string   `gorm:"-:all"`
	Roles            []Role   `gorm:"many2many:user_roles;"`
	Policies         []Policy `gorm:"many2many:user_policies;"`
	Active           bool
}

type Role struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Policy struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	Name string
}
type UserAccess struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	User       User
	UserId     uint
	AccessedAt time.Time
}
