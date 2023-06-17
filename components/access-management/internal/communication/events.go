package communication

import "time"

type SolutionSubmittedEvent struct {
	UserId       uint
	ProblemId    uint
	TestCode     string
	SolutionCode string
	Language     string
	CreatedAt    time.Time
}
