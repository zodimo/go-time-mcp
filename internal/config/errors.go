package config

import "fmt"

// ConfigError represents configuration-related errors
type ConfigError struct {
	Code    int    // Error code in range 3000-3999
	Message string // Human-readable error message
	Field   string // Configuration field that caused the error
	Err     error  // Underlying error if any
}

// Error implements the error interface
func (e *ConfigError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("config error [%d]: %s (field: %s)", e.Code, e.Message, e.Field)
	}
	return fmt.Sprintf("config error [%d]: %s", e.Code, e.Message)
}

// Unwrap returns the underlying error
func (e *ConfigError) Unwrap() error {
	return e.Err
}

// Configuration error codes (3000-3999 range)
const (
	ErrCodeInvalidMode     = 3001
	ErrCodeInvalidPort     = 3002
	ErrCodeInvalidTimeout  = 3003
	ErrCodeInvalidLogLevel = 3004
	ErrCodeParsingFailed   = 3005
)

// NewConfigError creates a new configuration error
func NewConfigError(code int, message, field string, err error) *ConfigError {
	return &ConfigError{
		Code:    code,
		Message: message,
		Field:   field,
		Err:     err,
	}
}

// NewInvalidModeError creates an error for invalid mode
func NewInvalidModeError(mode string) *ConfigError {
	return NewConfigError(
		ErrCodeInvalidMode,
		fmt.Sprintf("invalid mode '%s': must be 'sse' or 'stdio'", mode),
		"mode",
		nil,
	)
}

// NewInvalidPortError creates an error for invalid port
func NewInvalidPortError(port int) *ConfigError {
	return NewConfigError(
		ErrCodeInvalidPort,
		fmt.Sprintf("invalid port %d: must be between 1 and 65535", port),
		"port",
		nil,
	)
}

// NewInvalidTimeoutError creates an error for invalid timeout
func NewInvalidTimeoutError(timeout string) *ConfigError {
	return NewConfigError(
		ErrCodeInvalidTimeout,
		fmt.Sprintf("invalid timeout '%s': must be positive duration", timeout),
		"timeout",
		nil,
	)
}

// NewInvalidLogLevelError creates an error for invalid log level
func NewInvalidLogLevelError(level string) *ConfigError {
	return NewConfigError(
		ErrCodeInvalidLogLevel,
		fmt.Sprintf("invalid log level '%s': must be one of debug, info, warn, error", level),
		"log-level",
		nil,
	)
}
