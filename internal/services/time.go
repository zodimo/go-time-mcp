package services

import (
	"fmt"
	"strconv"
	"time"
)

// TimeService provides time-related operations
type TimeService interface {
	GetCurrentTime(timezone string) (time.Time, error)
	GetUnixTimestamp() int64
	FormatTime(t time.Time, format string) (string, error)
	ValidateTimezone(timezone string) error
	ValidateFormat(format string) error
}

// timeService implements TimeService interface
type timeService struct{}

// NewTimeService creates a new time service instance
func NewTimeService() TimeService {
	return &timeService{}
}

// GetCurrentTime returns the current time in the specified timezone
func (ts *timeService) GetCurrentTime(timezone string) (time.Time, error) {
	if timezone == "" {
		return time.Now().UTC(), nil
	}

	// Validate timezone first
	if err := ts.ValidateTimezone(timezone); err != nil {
		return time.Time{}, err
	}

	// Load the timezone
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, NewTimeServiceError(
			ErrCodeInvalidTimezone,
			fmt.Sprintf("failed to load timezone '%s'", timezone),
			"timezone",
			err,
		)
	}

	return time.Now().In(loc), nil
}

// GetUnixTimestamp returns the current Unix timestamp
func (ts *timeService) GetUnixTimestamp() int64 {
	return time.Now().Unix()
}

// FormatTime formats a time according to the specified format string
func (ts *timeService) FormatTime(t time.Time, format string) (string, error) {
	if format == "" {
		// Default to RFC3339 format
		return t.Format(time.RFC3339), nil
	}

	// Validate format first
	if err := ts.ValidateFormat(format); err != nil {
		return "", err
	}

	// Convert common format patterns to Go time format
	goFormat := convertToGoTimeFormat(format)

	return t.Format(goFormat), nil
}

// ValidateTimezone checks if a timezone string is valid
func (ts *timeService) ValidateTimezone(timezone string) error {
	if timezone == "" {
		return nil // Empty timezone is valid (defaults to UTC)
	}

	// Check for common timezone formats
	if isValidTimezoneFormat(timezone) {
		return nil
	}

	// Try to load the timezone to validate it
	_, err := time.LoadLocation(timezone)
	if err != nil {
		return NewInvalidTimezoneError(timezone, err)
	}

	return nil
}

// ValidateFormat checks if a format string is valid and safe
func (ts *timeService) ValidateFormat(format string) error {
	if format == "" {
		return nil // Empty format is valid (uses default)
	}

	// Check for potentially dangerous format strings
	if containsDangerousPatterns(format) {
		return NewInvalidFormatError(format, "format contains potentially dangerous patterns")
	}

	// Test the format with a known time to ensure it's valid
	testTime := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)
	goFormat := convertToGoTimeFormat(format)

	_, err := time.Parse(goFormat, testTime.Format(goFormat))
	if err != nil {
		return NewInvalidFormatError(format, "invalid time format")
	}

	return nil
}

// isValidTimezoneFormat checks for common timezone formats
func isValidTimezoneFormat(timezone string) bool {
	// Check for UTC variants
	if timezone == "UTC" || timezone == "GMT" {
		return true
	}

	// Check for offset format (+/-HH:MM or +/-HHMM)
	if len(timezone) >= 3 && (timezone[0] == '+' || timezone[0] == '-') {
		offset := timezone[1:]
		// Try to parse as offset
		if _, err := time.Parse("-07:00", timezone); err == nil {
			return true
		}
		if _, err := time.Parse("-0700", timezone); err == nil {
			return true
		}
		if _, err := strconv.Atoi(offset); err == nil && len(offset) <= 4 {
			return true
		}
	}

	// Check for common abbreviations (this is a basic check)
	commonTimezones := map[string]bool{
		"EST": true, "CST": true, "MST": true, "PST": true,
		"EDT": true, "CDT": true, "MDT": true, "PDT": true,
		"CET": true, "EET": true, "WET": true,
		"JST": true, "KST": true, "IST": true,
	}

	return commonTimezones[timezone]
}

// convertToGoTimeFormat converts common time format patterns to Go time format
func convertToGoTimeFormat(format string) string {
	// This is a simplified converter for common patterns
	// In a production system, you might want a more comprehensive converter

	// Common replacements
	replacements := map[string]string{
		"YYYY": "2006",
		"yyyy": "2006",
		"YY":   "06",
		"yy":   "06",
		"MM":   "01",
		"DD":   "02",
		"dd":   "02",
		"HH":   "15",
		"hh":   "03",
		"mm":   "04",
		"ss":   "05",
		"SSS":  "000",
	}

	result := format
	for pattern, replacement := range replacements {
		result = replaceAll(result, pattern, replacement)
	}

	return result
}

// replaceAll is a simple string replacement function
func replaceAll(s, old, new string) string {
	// Simple implementation - in production you might want to use strings.ReplaceAll
	// or a more sophisticated pattern matching
	result := ""
	i := 0
	for i < len(s) {
		if i+len(old) <= len(s) && s[i:i+len(old)] == old {
			result += new
			i += len(old)
		} else {
			result += string(s[i])
			i++
		}
	}
	return result
}

// containsDangerousPatterns checks for potentially dangerous format patterns
func containsDangerousPatterns(format string) bool {
	// Check for patterns that might be used for injection attacks
	dangerousPatterns := []string{
		"../", "./", "\\", "%", "$", "`", "$(", "${",
	}

	for _, pattern := range dangerousPatterns {
		if contains(format, pattern) {
			return true
		}
	}

	return false
}

// contains checks if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && indexOfSubstring(s, substr) >= 0
}

// indexOfSubstring finds the index of a substring in a string
func indexOfSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
