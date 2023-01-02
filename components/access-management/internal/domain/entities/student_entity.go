package entities

import (
	"time"
)

type StudentEntity struct {
	ID        int64
	Name      string
	Nickname  string
	DateBirth time.Time
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
