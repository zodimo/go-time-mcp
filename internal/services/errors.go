package services

import "fmt"

// TimeServiceError represents time service-related errors
type TimeServiceError struct {
	Code    int    // Error code in range 2000-2999
	Message string // Human-readable error message
	Field   string // Field that caused the error
	Err     error  // Underlying error if any
}

// Error implements the error interface
func (e *TimeServiceError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("time service error [%d]: %s (field: %s)", e.Code, e.Message, e.Field)
	}
	return fmt.Sprintf("time service error [%d]: %s", e.Code, e.Message)
}

// Unwrap returns the underlying error
func (e *TimeServiceError) Unwrap() error {
	return e.Err
}

// Time service error codes (2000-2999 range)
const (
	ErrCodeInvalidTimezone = 2001
	ErrCodeInvalidFormat   = 2002
	ErrCodeTimeOperation   = 2003
)

// NewTimeServiceError creates a new time service error
func NewTimeServiceError(code int, message, field string, err error) *TimeServiceError {
	return &TimeServiceError{
		Code:    code,
		Message: message,
		Field:   field,
		Err:     err,
	}
}

// NewInvalidTimezoneError creates an error for invalid timezone
func NewInvalidTimezoneError(timezone string, err error) *TimeServiceError {
	return NewTimeServiceError(
		ErrCodeInvalidTimezone,
		fmt.Sprintf("invalid timezone '%s'", timezone),
		"timezone",
		err,
	)
}

// NewInvalidFormatError creates an error for invalid format
func NewInvalidFormatError(format, reason string) *TimeServiceError {
	return NewTimeServiceError(
		ErrCodeInvalidFormat,
		fmt.Sprintf("invalid format '%s': %s", format, reason),
		"format",
		nil,
	)
}
