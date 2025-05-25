package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/zodimo/go-time-mcp/internal/config"
	"github.com/zodimo/go-time-mcp/internal/server"
	"github.com/zodimo/go-time-mcp/internal/services"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create time service
	timeService := services.NewTimeService()

	// Create MCP server
	mcpServer, err := server.NewServer(cfg, timeService)
	if err != nil {
		log.Fatalf("Failed to create MCP server: %v", err)
	}

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Shutdown signal received, stopping server...")
		cancel()
	}()

	// Start the server
	log.Printf("Starting go-time-mcp server in %s mode", cfg.Mode)
	if err := mcpServer.Start(ctx); err != nil {
		// Check if the error is due to context cancellation (graceful shutdown)
		if err == context.Canceled || err == context.DeadlineExceeded {
			log.Println("Server shutdown requested")
		} else {
			log.Fatalf("Server failed: %v", err)
		}
	}

	// Stop the server
	if err := mcpServer.Stop(); err != nil {
		log.Printf("Error stopping server: %v", err)
	}

	log.Println("Server stopped gracefully")
}
