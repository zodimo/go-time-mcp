package server

import (
	"context"
	"log"

	"github.com/mark3labs/mcp-go/server"
)

var _ Transport = (*stdioTransport)(nil)

// stdioTransport implements Transport for stdio mode
type stdioTransport struct{}

// NewStdioTransport creates a new stdio transport
func NewStdioTransport() Transport {
	return &stdioTransport{}
}

// Start starts the stdio transport
func (t *stdioTransport) Start(ctx context.Context, srv *server.MCPServer) error {
	log.Printf("Starting stdio transport")

	// Start the server in stdio mode
	return server.ServeStdio(srv)
}

// Stop stops the stdio transport
func (t *stdioTransport) Stop() error {
	log.Printf("Stopping stdio transport")
	// stdio transport stops when context is cancelled
	return nil
}

// Name returns the transport name
func (t *stdioTransport) Name() string {
	return "stdio"
}
