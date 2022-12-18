package entities

import "time"

type MediaEntity struct {
	ID        int64
	Filename  string
	MediaType string
	CreatedAt time.Time
	UpdatedAt time.Time
}
