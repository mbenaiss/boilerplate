package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/mbenaiss/boilerplate-go/api"
	"github.com/mbenaiss/boilerplate-go/db/sqlite"
	"github.com/mbenaiss/boilerplate-go/services"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	db := sqlite.NewClient(dbPath)
	service := services.NewService(db)
	server := api.NewServer(service)

	// Wait for an interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		err := server.Start()
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-quit

	server.Stop()
}
