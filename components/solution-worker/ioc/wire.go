package ioc

import (
	"solution-worker/config"
	"solution-worker/internal/communication"
	"solution-worker/internal/domain/handlers"
)

func InitSolutionSubmitHandler(outSolutionHandler *handlers.OutSolutionHandler) *handlers.SolutionSubmitHandler {
	handler := handlers.NewSolutionSubmitHandler(*outSolutionHandler)
	return handler
}

func InitOutSolutionHandler(config *config.Config) *handlers.OutSolutionHandler {
	publihser := ProvidePublisher(config.MessageBroker)
	handler := handlers.NewOutSolutionHandler(publihser)
	return &handler
}

func InitSubscribeHandler[T interface{}](config config.MessageBrokerConfig) communication.Subscriber[T] {
	subscribe := ProvideSubscribe[T](config)
	return subscribe
}
