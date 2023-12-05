package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"service"
)

func main() {
	// Load configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err.Error())
	}

	// Create new P2P service
	p2pService, err := service.NewP2PService(config)
	if err != nil {
		log.Fatalf("Failed to create P2P service: %s", err.Error())
	}

	// Start the P2P service
	go p2pService.Start()

	// Wait for termination signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	// Stop the P2P service
	p2pService.Stop()
}

func loadConfig() (*service.Config, error) {
	// Load configuration from config.go
	// This is just a placeholder, replace with your actual implementation
	return nil, nil
}
