package ui

import (
	"strings"
	"testing"
)

func TestColor(t *testing.T) {
	// Test with colors enabled
	ColorsEnabled = true
	result := Color(Red, "test")
	if !strings.Contains(result, "test") {
		t.Errorf("Color() should contain the text 'test', got %s", result)
	}
	if !strings.HasPrefix(result, Red) {
		t.Errorf("Color() should start with Red color code when colors enabled")
	}

	// Test with colors disabled
	ColorsEnabled = false
	result = Color(Red, "test")
	if result != "test" {
		t.Errorf("Color() with colors disabled should return plain text, got %s", result)
	}

	// Reset
	ColorsEnabled = true
}

func TestSuccessErrorWarningInfo(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string) string
	}{
		{"Success", Success},
		{"Error", Error},
		{"Warning", Warning},
		{"Info", Info},
		{"Highlight", Highlight},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn("test")
			if !strings.Contains(result, "test") {
				t.Errorf("%s() should contain 'test', got %s", tt.name, result)
			}
		})
	}
}

func TestProgressBar(t *testing.T) {
	tests := []struct {
		name     string
		progress float64
		width    int
	}{
		{"Zero progress", 0, 20},
		{"Half progress", 50, 20},
		{"Full progress", 100, 20},
		{"Over progress", 150, 20},
		{"Negative width", 50, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProgressBar(tt.progress, tt.width)
			if !strings.Contains(result, "%") {
				t.Errorf("ProgressBar() should contain '%%', got %s", result)
			}
			if !strings.Contains(result, "[") || !strings.Contains(result, "]") {
				t.Errorf("ProgressBar() should contain brackets, got %s", result)
			}
		})
	}
}

func TestBox(t *testing.T) {
	result := Box("Title", "Content line 1\nContent line 2")

	if !strings.Contains(result, "Title") {
		t.Errorf("Box() should contain title, got %s", result)
	}
	if !strings.Contains(result, "Content line 1") {
		t.Errorf("Box() should contain content, got %s", result)
	}
	if !strings.Contains(result, "+") {
		t.Errorf("Box() should contain border characters, got %s", result)
	}
}

func TestTable(t *testing.T) {
	headers := []string{"Name", "Age", "City"}
	rows := [][]string{
		{"Alice", "25", "NYC"},
		{"Bob", "30", "LA"},
	}

	result := Table(headers, rows)

	if !strings.Contains(result, "Name") {
		t.Errorf("Table() should contain header 'Name', got %s", result)
	}
	if !strings.Contains(result, "Alice") {
		t.Errorf("Table() should contain row data 'Alice', got %s", result)
	}

	// Test empty headers
	emptyResult := Table([]string{}, rows)
	if emptyResult != "" {
		t.Errorf("Table() with empty headers should return empty string, got %s", emptyResult)
	}
}

func TestSpinner(t *testing.T) {
	for i := 0; i < 10; i++ {
		result := Spinner(i)
		if result == "" {
			t.Errorf("Spinner(%d) should not be empty", i)
		}
	}
}

func TestBanner(t *testing.T) {
	result := Banner("1.0.0")
	if !strings.Contains(result, "DoS") {
		t.Errorf("Banner() should contain DoS, got %s", result)
	}
	if !strings.Contains(result, "1.0.0") {
		t.Errorf("Banner() should contain version, got %s", result)
	}
}

func TestFormatFunctions(t *testing.T) {
	ColorsEnabled = true

	method := FormatMethod("GET")
	if !strings.Contains(method, "GET") {
		t.Errorf("FormatMethod() should contain 'GET', got %s", method)
	}

	target := FormatTarget("example.com")
	if !strings.Contains(target, "example.com") {
		t.Errorf("FormatTarget() should contain 'example.com', got %s", target)
	}

	number := FormatNumber("1000")
	if !strings.Contains(number, "1000") {
		t.Errorf("FormatNumber() should contain '1000', got %s", number)
	}
}

func TestTitleHeader(t *testing.T) {
	ColorsEnabled = true

	title := Title("Test Title")
	if !strings.Contains(title, "Test Title") {
		t.Errorf("Title() should contain 'Test Title', got %s", title)
	}

	header := Header("Test Header")
	if !strings.Contains(header, "Test Header") {
		t.Errorf("Header() should contain 'Test Header', got %s", header)
	}
}
