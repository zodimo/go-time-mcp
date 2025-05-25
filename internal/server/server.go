package server

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/zodimo/go-time-mcp/internal/config"
	"github.com/zodimo/go-time-mcp/internal/services"
)

var _ Server = (*mcpServer)(nil)

// Server represents the MCP server
type Server interface {
	Start(ctx context.Context) error
	Stop() error
	RegisterToolHandler(name string, handler server.ToolHandlerFunc) error
}

// mcpServer implements Server interface
type mcpServer struct {
	config      *config.Config
	timeService services.TimeService
	server      *server.MCPServer
	transport   Transport
}

// Transport represents the transport layer (SSE or stdio)
type Transport interface {
	Start(ctx context.Context, srv *server.MCPServer) error
	Stop() error
	Name() string
}

// NewServer creates a new MCP server instance
func NewServer(cfg *config.Config, timeService services.TimeService) (Server, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}
	if timeService == nil {
		return nil, fmt.Errorf("time service cannot be nil")
	}

	// Create the MCP server
	mcpSrv := server.NewMCPServer("go-time-mcp", "1.0.0")

	// Create transport based on mode
	transport, err := createTransport(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create transport: %w", err)
	}

	srv := &mcpServer{
		config:      cfg,
		timeService: timeService,
		server:      mcpSrv,
		transport:   transport,
	}

	// Register tool handlers
	if err := srv.registerToolHandlers(); err != nil {
		return nil, fmt.Errorf("failed to register tool handlers: %w", err)
	}

	return srv, nil
}

// Start starts the MCP server
func (s *mcpServer) Start(ctx context.Context) error {
	log.Printf("Starting MCP server in %s mode", s.config.Mode)

	return s.transport.Start(ctx, s.server)
}

// Stop stops the MCP server
func (s *mcpServer) Stop() error {
	log.Printf("Stopping MCP server")

	return s.transport.Stop()
}

// RegisterToolHandler registers a tool handler with the MCP server
func (s *mcpServer) RegisterToolHandler(name string, handler server.ToolHandlerFunc) error {
	// Create tool definition
	tool := mcp.Tool{
		Name: name,
	}

	// Register the tool and handler
	s.server.AddTool(tool, handler)
	return nil
}

// registerToolHandlers registers all time-related tool handlers
func (s *mcpServer) registerToolHandlers() error {
	// Register getCurrentTime tool handler
	getCurrentTimeHandler := func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Extract parameters using mcp helper functions
		timezone := mcp.ParseString(request, "timezone", "")
		format := mcp.ParseString(request, "format", "")

		// Get current time
		currentTime, err := s.timeService.GetCurrentTime(timezone)
		if err != nil {
			return nil, err
		}

		// Format time
		formattedTime, err := s.timeService.FormatTime(currentTime, format)
		if err != nil {
			return nil, err
		}

		return mcp.NewToolResultText(formattedTime), nil
	}

	getCurrentTimeTool := mcp.Tool{
		Name:        "getCurrentTime",
		Description: "Get LIVE current time from system clock. IMPORTANT FOR LLMs: Use this tool whenever you need current time information, as LLMs cannot determine the current time independently due to training data cutoffs. This tool provides real-time data directly from the system clock with timezone support and custom formatting options.",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"timezone": map[string]interface{}{
					"type":        "string",
					"description": "Timezone (IANA format, e.g., 'America/New_York', or empty for UTC)",
				},
				"format": map[string]interface{}{
					"type":        "string",
					"description": "Time format string (optional, defaults to RFC3339)",
				},
			},
		},
	}

	s.server.AddTool(getCurrentTimeTool, getCurrentTimeHandler)

	// Register getUnixTimestamp tool handler
	getUnixTimestampHandler := func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		timestamp := s.timeService.GetUnixTimestamp()

		return mcp.NewToolResultText(fmt.Sprintf("%d", timestamp)), nil
	}

	getUnixTimestampTool := mcp.Tool{
		Name:        "getUnixTimestamp",
		Description: "Get LIVE current Unix timestamp from system clock. IMPORTANT FOR LLMs: Use this tool to obtain the precise current Unix timestamp, as LLMs cannot access real-time system data independently. This provides the current number of seconds since January 1, 1970 UTC, essential for time-based calculations and accurate timestamps.",
		InputSchema: mcp.ToolInputSchema{
			Type:       "object",
			Properties: map[string]interface{}{},
		},
	}

	s.server.AddTool(getUnixTimestampTool, getUnixTimestampHandler)

	log.Printf("Registered %d tools", 2)
	return nil
}

// createTransport creates the appropriate transport based on configuration
func createTransport(cfg *config.Config) (Transport, error) {
	switch cfg.Mode {
	case "sse":
		return NewSSETransport(cfg.Port), nil
	case "stdio":
		return NewStdioTransport(), nil
	default:
		return nil, fmt.Errorf("unsupported mode: %s", cfg.Mode)
	}
}
