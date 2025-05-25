package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config holds all configuration for the MCP server
type Config struct {
	Mode     string        // SSE or stdio
	Port     int           // Port for SSE mode
	Timeout  time.Duration // Request timeout
	LogLevel string        // Log level (debug, info, warn, error)
}

// Load parses command line flags and environment variables to create configuration
func Load() (*Config, error) {
	cfg := &Config{}

	// Define command line flags with defaults
	mode := flag.String("mode", getEnvOrDefault("MCP_MODE", "stdio"), "Server mode: 'sse' or 'stdio'")
	port := flag.Int("port", getEnvIntOrDefault("MCP_PORT", 8080), "Port for SSE mode")
	timeout := flag.Duration("timeout", getEnvDurationOrDefault("MCP_TIMEOUT", 30*time.Second), "Request timeout")
	logLevel := flag.String("log-level", getEnvOrDefault("MCP_LOG_LEVEL", "info"), "Log level: debug, info, warn, error")

	// Parse command line flags
	flag.Parse()

	// Set configuration values
	cfg.Mode = *mode
	cfg.Port = *port
	cfg.Timeout = *timeout
	cfg.LogLevel = *logLevel

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return cfg, nil
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	// Validate mode
	if c.Mode != "sse" && c.Mode != "stdio" {
		return NewInvalidModeError(c.Mode)
	}

	// Validate port for SSE mode
	if c.Mode == "sse" {
		if c.Port < 1 || c.Port > 65535 {
			return NewInvalidPortError(c.Port)
		}
	}

	// Validate timeout
	if c.Timeout <= 0 {
		return NewInvalidTimeoutError(c.Timeout.String())
	}

	// Validate log level
	validLogLevels := map[string]bool{
		"debug": true,
		"info":  true,
		"warn":  true,
		"error": true,
	}
	if !validLogLevels[c.LogLevel] {
		return NewInvalidLogLevelError(c.LogLevel)
	}

	return nil
}

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvIntOrDefault returns environment variable as int or default if not set/invalid
func getEnvIntOrDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvDurationOrDefault returns environment variable as duration or default if not set/invalid
func getEnvDurationOrDefault(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
