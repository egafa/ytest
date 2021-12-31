package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	handler "github.com/egafa/ytest/api/handler"
	model "github.com/egafa/ytest/api/model"
)

func main() {

	mm := model.MapMetric{}
	mm.GaugeData = make(map[string]float64)
	mm.CounterData = make(map[string][]int64)

	addr := "127.0.0.1:8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/update/", handler.MetricHandler(mm))

	srv := &http.Server{
		Handler: mux,
	}

	srv.Addr = addr

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed

}
