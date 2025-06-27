package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Health struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if err := InitPgxPool(ctx, cfg); err != nil {
		_ = fmt.Errorf("Failed to initialize PostgreSQL pool: %v\nDatabase tests will fail and/or may crash the program", err)
	}
	defer ClosePgxPool()

	if err := InitRedis(ctx, cfg); err != nil {
		_ = fmt.Errorf("Failed to initialize Redis: %v\n Redis tests will fail and/or may crash the program", err)
	}
	defer CloseRedis()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		health := Health{
			Status:    "healthy",
			Timestamp: time.Now().UTC(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(health)
	})

	// json serialization
	r.Get("/user/json", HandleUserSerialization)

	// db tests
	r.Get("/user/db/{id}", HandleDbReadTest)

	r.Post("/user/db", HandleDbWriteTest)

	// cache/fallback tests
	r.Get("/user/cache/{id}", HandleCacheUser)

	// concurrency tests with simulated delay
	r.Get("/user/concurrent", HandleConcurrent)

	server := &http.Server{
		Addr:    ":" + "8080",
		Handler: r,
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutting down server...")
		cancel() // signal context cancellation

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}
	}()

	log.Printf("Server starting on port %s\n", "8080")
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("Server exited gracefully")
}
