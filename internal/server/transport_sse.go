package server

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/server"
)

var _ Transport = (*sseTransport)(nil)

// sseTransport implements Transport for SSE mode
type sseTransport struct {
	port      int
	sseServer *server.SSEServer
}

// NewSSETransport creates a new SSE transport
func NewSSETransport(port int) Transport {
	return &sseTransport{
		port: port,
	}
}

// Start starts the SSE transport
func (t *sseTransport) Start(ctx context.Context, srv *server.MCPServer) error {
	log.Printf("Starting SSE transport on port %d", t.port)

	// Create SSE server
	t.sseServer = server.NewSSEServer(srv)

	// Start the SSE server in a goroutine to make it non-blocking
	addr := fmt.Sprintf(":%d", t.port)
	errChan := make(chan error, 1)

	go func() {
		if err := t.sseServer.Start(addr); err != nil {
			errChan <- err
		}
	}()

	// Wait for either context cancellation or server error
	select {
	case <-ctx.Done():
		log.Printf("SSE transport context cancelled, shutting down")
		// Shutdown the server when context is cancelled
		if shutdownErr := t.sseServer.Shutdown(context.Background()); shutdownErr != nil {
			log.Printf("Error during SSE server shutdown: %v", shutdownErr)
		}
		return ctx.Err()
	case err := <-errChan:
		return fmt.Errorf("SSE server failed to start: %w", err)
	}
}

// Stop stops the SSE transport
func (t *sseTransport) Stop() error {
	log.Printf("Stopping SSE transport")

	if t.sseServer != nil {
		return t.sseServer.Shutdown(context.Background())
	}
	return nil
}

// Name returns the transport name
func (t *sseTransport) Name() string {
	return "sse"
}
