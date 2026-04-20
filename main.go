package main

import (
	"log"
	"os"
	"os/signal"

	"martinshaw.co/cloudduplicatefilefinder/config"
	"martinshaw.co/cloudduplicatefilefinder/server"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	config := config.LoadConfig()

	server := server.LoadServer(config)
	server.Announce()

	<-quit
	log.Printf("Shutting down server...")

	if err := server.App.Shutdown(); err != nil {
		log.Fatalf("Server Shutdown Failed: %v", err)
	}
}
