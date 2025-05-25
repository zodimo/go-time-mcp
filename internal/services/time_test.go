package services

import (
	"testing"
	"time"
)

func TestTimeService_GetUnixTimestamp(t *testing.T) {
	ts := NewTimeService()

	timestamp := ts.GetUnixTimestamp()

	// Verify timestamp is reasonable (within last minute)
	now := time.Now().Unix()
	if timestamp < now-60 || timestamp > now+60 {
		t.Errorf("Timestamp %d is not within reasonable range of %d", timestamp, now)
	}
}

func TestTimeService_GetCurrentTime(t *testing.T) {
	ts := NewTimeService()

	tests := []struct {
		name     string
		timezone string
		wantErr  bool
	}{
		{
			name:     "UTC timezone",
			timezone: "UTC",
			wantErr:  false,
		},
		{
			name:     "Empty timezone (defaults to UTC)",
			timezone: "",
			wantErr:  false,
		},
		{
			name:     "America/New_York timezone",
			timezone: "America/New_York",
			wantErr:  false,
		},
		{
			name:     "Invalid timezone",
			timezone: "Invalid/Timezone",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currentTime, err := ts.GetCurrentTime(tt.timezone)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error for timezone %s, but got none", tt.timezone)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error for timezone %s: %v", tt.timezone, err)
				return
			}

			// Verify time is reasonable (within last minute)
			now := time.Now()
			diff := now.Sub(currentTime)
			if diff < -time.Minute || diff > time.Minute {
				t.Errorf("Time %v is not within reasonable range of %v", currentTime, now)
			}
		})
	}
}

func TestTimeService_FormatTime(t *testing.T) {
	ts := NewTimeService()
	testTime := time.Date(2024, 1, 15, 14, 30, 45, 0, time.UTC)

	tests := []struct {
		name     string
		format   string
		expected string
		wantErr  bool
	}{
		{
			name:     "Empty format (defaults to RFC3339)",
			format:   "",
			expected: "2024-01-15T14:30:45Z",
			wantErr:  false,
		},
		{
			name:     "YYYY-MM-DD format",
			format:   "YYYY-MM-DD",
			expected: "2024-01-15",
			wantErr:  false,
		},
		{
			name:     "Dangerous format",
			format:   "../etc/passwd",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ts.FormatTime(testTime, tt.format)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error for format %s, but got none", tt.format)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error for format %s: %v", tt.format, err)
				return
			}

			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestTimeService_ValidateTimezone(t *testing.T) {
	ts := NewTimeService()

	tests := []struct {
		name     string
		timezone string
		wantErr  bool
	}{
		{
			name:     "Valid UTC",
			timezone: "UTC",
			wantErr:  false,
		},
		{
			name:     "Valid GMT",
			timezone: "GMT",
			wantErr:  false,
		},
		{
			name:     "Valid IANA timezone",
			timezone: "America/New_York",
			wantErr:  false,
		},
		{
			name:     "Valid offset +05:00",
			timezone: "+05:00",
			wantErr:  false,
		},
		{
			name:     "Valid offset -0800",
			timezone: "-0800",
			wantErr:  false,
		},
		{
			name:     "Empty timezone",
			timezone: "",
			wantErr:  false,
		},
		{
			name:     "Invalid timezone",
			timezone: "Invalid/Timezone",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ts.ValidateTimezone(tt.timezone)

			if tt.wantErr && err == nil {
				t.Errorf("Expected error for timezone %s, but got none", tt.timezone)
			}

			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error for timezone %s: %v", tt.timezone, err)
			}
		})
	}
}

func TestTimeService_ValidateFormat(t *testing.T) {
	ts := NewTimeService()

	tests := []struct {
		name    string
		format  string
		wantErr bool
	}{
		{
			name:    "Empty format",
			format:  "",
			wantErr: false,
		},
		{
			name:    "Valid YYYY-MM-DD format",
			format:  "YYYY-MM-DD",
			wantErr: false,
		},
		{
			name:    "Dangerous format with ../",
			format:  "../etc/passwd",
			wantErr: true,
		},
		{
			name:    "Dangerous format with $",
			format:  "$(rm -rf /)",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ts.ValidateFormat(tt.format)

			if tt.wantErr && err == nil {
				t.Errorf("Expected error for format %s, but got none", tt.format)
			}

			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error for format %s: %v", tt.format, err)
			}
		})
	}
}
