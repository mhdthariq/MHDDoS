package ui

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
)

// ValidationResult holds the result of a validation
type ValidationResult struct {
	Valid   bool
	Message string
	Hint    string
}

// ValidateURL validates a URL and returns helpful feedback
func ValidateURL(target string) ValidationResult {
	// Add http:// if no scheme is present
	normalizedTarget := target
	if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
		normalizedTarget = "http://" + target
	}

	u, err := url.Parse(normalizedTarget)
	if err != nil {
		return ValidationResult{
			Valid:   false,
			Message: fmt.Sprintf("Invalid URL format: %v", err),
			Hint:    "Example: http://example.com or https://example.com",
		}
	}

	if u.Host == "" {
		return ValidationResult{
			Valid:   false,
			Message: "URL is missing host",
			Hint:    "Example: http://example.com or https://example.com/path",
		}
	}

	return ValidationResult{
		Valid:   true,
		Message: "URL is valid",
	}
}

// ValidateHostPort validates a host:port combination
func ValidateHostPort(target string) ValidationResult {
	// Remove any protocol prefix
	cleanTarget := target
	if strings.Contains(target, "://") {
		parts := strings.SplitN(target, "://", 2)
		if len(parts) > 1 {
			cleanTarget = parts[1]
		}
	}

	// Check if it contains a port
	if !strings.Contains(cleanTarget, ":") {
		return ValidationResult{
			Valid:   false,
			Message: "Missing port number",
			Hint:    "Example: 192.168.1.1:80 or example.com:443",
		}
	}

	host, portStr, err := net.SplitHostPort(cleanTarget)
	if err != nil {
		return ValidationResult{
			Valid:   false,
			Message: fmt.Sprintf("Invalid host:port format: %v", err),
			Hint:    "Example: 192.168.1.1:80 or example.com:443",
		}
	}

	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return ValidationResult{
			Valid:   false,
			Message: "Invalid port number (must be 1-65535)",
			Hint:    "Common ports: 80 (HTTP), 443 (HTTPS), 25565 (Minecraft)",
		}
	}

	// Try to validate the host
	if net.ParseIP(host) == nil {
		// Not an IP, check if it's a valid hostname
		if len(host) == 0 {
			return ValidationResult{
				Valid:   false,
				Message: "Empty hostname",
				Hint:    "Example: 192.168.1.1:80 or example.com:443",
			}
		}
	}

	return ValidationResult{
		Valid:   true,
		Message: "Host and port are valid",
	}
}

// ValidateThreads validates thread count
func ValidateThreads(threads int) ValidationResult {
	if threads < 1 {
		return ValidationResult{
			Valid:   false,
			Message: "Thread count must be at least 1",
			Hint:    "Recommended: 100-500 threads for most attacks",
		}
	}

	if threads > 10000 {
		return ValidationResult{
			Valid:   true,
			Message: fmt.Sprintf("Warning: %d threads is very high", threads),
			Hint:    "High thread counts may consume too many resources. Consider 100-1000 threads.",
		}
	}

	return ValidationResult{
		Valid:   true,
		Message: "Thread count is valid",
	}
}

// ValidateDuration validates attack duration
func ValidateDuration(duration int) ValidationResult {
	if duration < 1 {
		return ValidationResult{
			Valid:   false,
			Message: "Duration must be at least 1 second",
			Hint:    "Example: 60 for 1 minute, 300 for 5 minutes",
		}
	}

	if duration > 86400 {
		return ValidationResult{
			Valid:   true,
			Message: "Warning: Duration exceeds 24 hours",
			Hint:    "Very long attacks may trigger detection systems.",
		}
	}

	return ValidationResult{
		Valid:   true,
		Message: "Duration is valid",
	}
}

// ValidateRPC validates requests per connection
func ValidateRPC(rpc int) ValidationResult {
	if rpc < 1 {
		return ValidationResult{
			Valid:   false,
			Message: "RPC must be at least 1",
			Hint:    "Recommended: 50-100 for most Layer 7 attacks",
		}
	}

	if rpc > 10000 {
		return ValidationResult{
			Valid:   true,
			Message: "Warning: RPC is very high",
			Hint:    "High RPC values may cause connection timeouts. Consider 50-500.",
		}
	}

	return ValidationResult{
		Valid:   true,
		Message: "RPC is valid",
	}
}

// ValidateProxyType validates proxy type
func ValidateProxyType(proxyType int) ValidationResult {
	validTypes := map[int]string{
		0: "All (mixed)",
		1: "HTTP",
		4: "SOCKS4",
		5: "SOCKS5",
		6: "Random",
	}

	if typeName, ok := validTypes[proxyType]; ok {
		return ValidationResult{
			Valid:   true,
			Message: fmt.Sprintf("Proxy type: %s", typeName),
		}
	}

	return ValidationResult{
		Valid:   false,
		Message: fmt.Sprintf("Invalid proxy type: %d", proxyType),
		Hint:    "Valid types: 0 (all), 1 (HTTP), 4 (SOCKS4), 5 (SOCKS5), 6 (random)",
	}
}

// SuggestMethod suggests similar methods based on input
func SuggestMethod(input string, validMethods []string) []string {
	input = strings.ToUpper(input)
	var suggestions []string

	for _, method := range validMethods {
		// Check for prefix match
		if strings.HasPrefix(method, input) {
			suggestions = append(suggestions, method)
			continue
		}

		// Check for substring match
		if len(input) >= 2 && strings.Contains(method, input) {
			suggestions = append(suggestions, method)
			continue
		}

		// Simple Levenshtein-like check for close matches
		if levenshteinClose(input, method) {
			suggestions = append(suggestions, method)
		}
	}

	// Limit to 5 suggestions
	if len(suggestions) > 5 {
		suggestions = suggestions[:5]
	}

	return suggestions
}

// levenshteinClose checks if two strings are within 2 edits of each other
func levenshteinClose(a, b string) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}

	// Quick length check
	if abs(len(a)-len(b)) > 2 {
		return false
	}

	// Count matching characters
	matches := 0
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	for i := 0; i < minLen; i++ {
		if a[i] == b[i] {
			matches++
		}
	}

	// If more than 50% match, consider it close
	return float64(matches)/float64(max(len(a), len(b))) > 0.5
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// FormatValidationError formats a validation result as a user-friendly error
func FormatValidationError(result ValidationResult, fieldName string) string {
	var sb strings.Builder

	sb.WriteString(Error(fmt.Sprintf("Invalid %s: ", fieldName)))
	sb.WriteString(result.Message)

	if result.Hint != "" {
		sb.WriteString("\n")
		sb.WriteString(Info("  Hint: "))
		sb.WriteString(result.Hint)
	}

	return sb.String()
}
