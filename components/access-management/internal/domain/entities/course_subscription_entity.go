package entities

import "time"

type Subscription struct {
	ID             int64
	StudentId      int64
	Student        StudentEntity
	CourseId       int64
	Course         CourseEntity
	SubscriptionAt time.Time
}
