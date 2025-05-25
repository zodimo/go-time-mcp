package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"go-time-mcp",
		"1.0.0",
		"MCP server for time services with timezone and format support",
	)

	// TODO: Register time-related tools
	// - getCurrentTime tool
	// - getUnixTimestamp tool
	// - formatTime tool

	// TODO: Start server in appropriate mode (SSE or stdio)
	// Based on command line arguments or environment variables

	ctx := context.Background()

	// Placeholder - will be implemented based on mode selection
	fmt.Println("go-time-mcp server starting...")

	// TODO: Implement actual server startup logic
	log.Printf("Server initialized at %s", time.Now().Format(time.RFC3339))

	// Keep server running
	select {
	case <-ctx.Done():
		log.Println("Server shutting down...")
	}
}
