package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	h "github.com/ckang21/go-worker-service/internal/http"
	"github.com/ckang21/go-worker-service/internal/queue"
	"github.com/ckang21/go-worker-service/internal/tasks"
)

func main() {
	addr := getenv("ADDR", ":8080")

	// In-memory store + worker queue
	store := tasks.NewStore()
	q := queue.New[tasks.Job](128)
	worker := tasks.NewWorker(q, store, log.Default())
	go worker.Run()

	mux := http.NewServeMux()
	h.RegisterRoutes(mux, store, q)

	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	go func() {
		log.Printf("server listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	worker.Stop()
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
