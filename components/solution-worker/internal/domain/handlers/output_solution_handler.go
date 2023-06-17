package handlers

import (
	"github.com/google/uuid"
	"solution-worker/internal/communication"
	"solution-worker/internal/domain/events"
)

type OutSolutionHandler struct {
	publisher communication.Publisher
}

func NewOutSolutionHandler(pubisher communication.Publisher) OutSolutionHandler {
	return OutSolutionHandler{
		publisher: pubisher,
	}
}

func (h OutSolutionHandler) Handler(event events.OutputSolutionEvent) error {
	err := h.publisher.Send("output-solution", event, uuid.New().String())
	if err != nil {
		return err
	}
	return nil
}
