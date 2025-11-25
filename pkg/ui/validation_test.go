package ui

import (
	"testing"
)

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name    string
		target  string
		isValid bool
	}{
		{"Valid HTTP URL", "http://example.com", true},
		{"Valid HTTPS URL", "https://example.com", true},
		{"URL without scheme", "example.com", true},
		{"URL with path", "http://example.com/path", true},
		{"Empty URL", "", false},
		{"Just protocol", "http://", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateURL(tt.target)
			if result.Valid != tt.isValid {
				t.Errorf("ValidateURL(%s).Valid = %v; want %v", tt.target, result.Valid, tt.isValid)
			}
		})
	}
}

func TestValidateHostPort(t *testing.T) {
	tests := []struct {
		name    string
		target  string
		isValid bool
	}{
		{"Valid IP:port", "192.168.1.1:80", true},
		{"Valid hostname:port", "example.com:443", true},
		{"Missing port", "192.168.1.1", false},
		{"Invalid port string", "192.168.1.1:abc", false},
		{"Port too high", "192.168.1.1:99999", false},
		{"With protocol prefix", "http://192.168.1.1:80", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateHostPort(tt.target)
			if result.Valid != tt.isValid {
				t.Errorf("ValidateHostPort(%s).Valid = %v; want %v, message: %s", tt.target, result.Valid, tt.isValid, result.Message)
			}
		})
	}
}

func TestValidateThreads(t *testing.T) {
	tests := []struct {
		name    string
		threads int
		isValid bool
	}{
		{"Valid threads", 100, true},
		{"Zero threads", 0, false},
		{"Negative threads", -1, false},
		{"High threads warning", 15000, true},
		{"One thread", 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateThreads(tt.threads)
			if result.Valid != tt.isValid {
				t.Errorf("ValidateThreads(%d).Valid = %v; want %v", tt.threads, result.Valid, tt.isValid)
			}
		})
	}
}

func TestValidateDuration(t *testing.T) {
	tests := []struct {
		name     string
		duration int
		isValid  bool
	}{
		{"Valid duration", 60, true},
		{"Zero duration", 0, false},
		{"Negative duration", -1, false},
		{"Very long duration", 100000, true},
		{"One second", 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateDuration(tt.duration)
			if result.Valid != tt.isValid {
				t.Errorf("ValidateDuration(%d).Valid = %v; want %v", tt.duration, result.Valid, tt.isValid)
			}
		})
	}
}

func TestValidateRPC(t *testing.T) {
	tests := []struct {
		name    string
		rpc     int
		isValid bool
	}{
		{"Valid RPC", 100, true},
		{"Zero RPC", 0, false},
		{"Negative RPC", -1, false},
		{"High RPC warning", 15000, true},
		{"One RPC", 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateRPC(tt.rpc)
			if result.Valid != tt.isValid {
				t.Errorf("ValidateRPC(%d).Valid = %v; want %v", tt.rpc, result.Valid, tt.isValid)
			}
		})
	}
}

func TestValidateProxyType(t *testing.T) {
	tests := []struct {
		name      string
		proxyType int
		isValid   bool
	}{
		{"HTTP proxy", 1, true},
		{"SOCKS4 proxy", 4, true},
		{"SOCKS5 proxy", 5, true},
		{"All proxies", 0, true},
		{"Random", 6, true},
		{"Invalid type", 99, false},
		{"Negative type", -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateProxyType(tt.proxyType)
			if result.Valid != tt.isValid {
				t.Errorf("ValidateProxyType(%d).Valid = %v; want %v", tt.proxyType, result.Valid, tt.isValid)
			}
		})
	}
}

func TestSuggestMethod(t *testing.T) {
	validMethods := []string{"GET", "POST", "HEAD", "GSB", "TCP", "UDP"}

	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"Prefix match", "G", []string{"GET", "GSB"}},
		{"Exact match partial", "GE", []string{"GET"}},
		{"No match", "XYZ", []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SuggestMethod(tt.input, validMethods)
			if len(result) != len(tt.expected) {
				t.Errorf("SuggestMethod(%s) returned %d suggestions; want %d", tt.input, len(result), len(tt.expected))
			}
		})
	}
}

func TestFormatValidationError(t *testing.T) {
	result := ValidationResult{
		Valid:   false,
		Message: "Test error",
		Hint:    "Test hint",
	}

	output := FormatValidationError(result, "test field")

	if output == "" {
		t.Error("FormatValidationError() should not return empty string")
	}
}
