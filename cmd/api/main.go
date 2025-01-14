package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/minhaz11/crud/internal/config"
	student "github.com/minhaz11/crud/internal/http/handlers"
	"github.com/minhaz11/crud/internal/storage/sqlite"
)

func main() {
	//load config

	cfg := config.MustLoad()

	//database setup

	storage, err := sqlite.New(cfg)
	
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Sotrage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0"))

	//setup router

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetStudent(storage))
	router.HandleFunc("GET /api/students", student.GetStudentList(storage))

	//setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	// gracefully shutdown server
	
	slog.Info("shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")

}
