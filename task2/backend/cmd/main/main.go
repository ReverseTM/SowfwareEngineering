package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"software-engineering-2/internal/app"
	"syscall"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	application := app.NewApp()

	go func() {
		err := application.Run()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	<-signalChan
	err := application.Stop()
	if err != nil {
		log.Fatal(err)
	}
}
