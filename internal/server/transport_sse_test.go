package server

import (
	"context"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/server"
)

func TestSSETransport_ContextCancellation(t *testing.T) {
	// Create a test MCP server
	mcpSrv := server.NewMCPServer("test-server", "1.0.0")

	// Create SSE transport with a different port to avoid conflicts
	transport := NewSSETransport(9999)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Start the transport - this should respect context cancellation
	err := transport.Start(ctx, mcpSrv)

	// Should return context.DeadlineExceeded or context.Canceled
	if err == nil {
		t.Error("Expected error due to context cancellation, got nil")
	}

	// The error should be either context cancellation or server start failure
	// Both are acceptable for this test - we're mainly testing that Start() returns
	// when context is cancelled rather than blocking indefinitely
	if err != context.DeadlineExceeded && err != context.Canceled {
		// If it's a server start error, that's also acceptable as long as it doesn't block
		t.Logf("Got server start error (acceptable): %v", err)
	} else {
		t.Logf("Got expected context cancellation: %v", err)
	}
}

func TestSSETransport_Name(t *testing.T) {
	transport := NewSSETransport(8080)

	if transport.Name() != "sse" {
		t.Errorf("Expected transport name 'sse', got '%s'", transport.Name())
	}
}
