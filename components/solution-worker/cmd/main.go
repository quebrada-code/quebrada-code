package main

import (
	"fmt"
	"os"
	"os/signal"
	"solution-worker/config"
	"solution-worker/internal/domain/events"
	"solution-worker/ioc"
	"syscall"
)

// TODO: Adicionar Health Check

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	cfg, err := config.Init()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed get app configuration: %s\n", err)
		os.Exit(1)
		return
	}

	outSolutionHandler := ioc.InitOutSolutionHandler(cfg)
	solutionHandler := ioc.InitSolutionSubmitHandler(outSolutionHandler)
	solutionSubmitSubscribe := ioc.InitSubscribeHandler[events.SolutionSubmittedEvent](cfg.MessageBroker)

	select {
	case sig := <-sigchan:
		fmt.Printf("Caught signal %v: terminating\n", sig)
	default:
		err = solutionSubmitSubscribe.Subscribe("solution-submitted", solutionHandler.Handler)
		if err != nil {
			return
		}
	}

}
