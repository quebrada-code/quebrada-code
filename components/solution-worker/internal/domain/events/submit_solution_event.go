package events

import (
	"encoding/base64"
	"time"
)

type SolutionSubmittedEvent struct {
	UserId       uint
	ProblemId    uint
	TestCode     string
	SolutionCode string
	Language     string
	CreatedAt    time.Time
}

func (e SolutionSubmittedEvent) GetSolutionCodeDecoded() string {
	return decode64(e.SolutionCode)
}

func (e SolutionSubmittedEvent) GetTestCodeDecoded() string {
	return decode64(e.TestCode)
}

func decode64(text string) string {
	solutionCode, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		panic(err)
	}
	return string(solutionCode[:])
}
