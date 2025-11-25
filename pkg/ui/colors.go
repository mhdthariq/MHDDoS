package ui

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ANSI color codes
const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	Dim   = "\033[2m"

	// Foreground colors
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	// Bright foreground colors
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"
)

// ColorsEnabled determines if ANSI colors should be used
var ColorsEnabled = true

func init() {
	// Disable colors on Windows unless using a modern terminal
	if runtime.GOOS == "windows" {
		// Check if running in a modern terminal that supports ANSI
		if os.Getenv("WT_SESSION") == "" && os.Getenv("TERM_PROGRAM") == "" {
			ColorsEnabled = false
		}
	}

	// Respect NO_COLOR environment variable
	if os.Getenv("NO_COLOR") != "" {
		ColorsEnabled = false
	}
}

// Color applies color to a string if colors are enabled
func Color(color, text string) string {
	if !ColorsEnabled {
		return text
	}
	return color + text + Reset
}

// Success formats text as success (green)
func Success(text string) string {
	return Color(Green, text)
}

// Error formats text as error (red)
func Error(text string) string {
	return Color(Red, text)
}

// Warning formats text as warning (yellow)
func Warning(text string) string {
	return Color(Yellow, text)
}

// Info formats text as info (cyan)
func Info(text string) string {
	return Color(Cyan, text)
}

// Highlight formats text as highlighted (magenta)
func Highlight(text string) string {
	return Color(Magenta, text)
}

// Title formats text as title (bold blue)
func Title(text string) string {
	if !ColorsEnabled {
		return text
	}
	return Bold + Blue + text + Reset
}

// Header formats text as header (bold cyan)
func Header(text string) string {
	if !ColorsEnabled {
		return text
	}
	return Bold + Cyan + text + Reset
}

// PrintSuccess prints a success message
func PrintSuccess(format string, args ...interface{}) {
	fmt.Print(Success("[OK] "))
	fmt.Printf(format, args...)
	fmt.Println()
}

// PrintError prints an error message
func PrintError(format string, args ...interface{}) {
	fmt.Print(Error("[ERROR] "))
	fmt.Printf(format, args...)
	fmt.Println()
}

// PrintWarning prints a warning message
func PrintWarning(format string, args ...interface{}) {
	fmt.Print(Warning("[WARN] "))
	fmt.Printf(format, args...)
	fmt.Println()
}

// PrintInfo prints an info message
func PrintInfo(format string, args ...interface{}) {
	fmt.Print(Info("[INFO] "))
	fmt.Printf(format, args...)
	fmt.Println()
}

// ProgressBar creates a simple progress bar
func ProgressBar(progress float64, width int) string {
	if width <= 0 {
		width = 30
	}

	filled := int(progress / 100 * float64(width))
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}

	bar := strings.Repeat("=", filled) + strings.Repeat("-", width-filled)

	if ColorsEnabled {
		var color string
		if progress < 33 {
			color = Red
		} else if progress < 66 {
			color = Yellow
		} else {
			color = Green
		}
		return fmt.Sprintf("[%s%s%s] %.1f%%", color, bar, Reset, progress)
	}

	return fmt.Sprintf("[%s] %.1f%%", bar, progress)
}

// Box creates a box around text
func Box(title, content string) string {
	lines := strings.Split(content, "\n")

	// Find max width
	maxWidth := len(title)
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// Build box
	var sb strings.Builder
	topBorder := "+" + strings.Repeat("-", maxWidth+2) + "+"
	bottomBorder := "+" + strings.Repeat("-", maxWidth+2) + "+"

	if ColorsEnabled {
		sb.WriteString(Cyan)
	}

	sb.WriteString(topBorder + "\n")

	if title != "" {
		padding := maxWidth - len(title)
		leftPad := padding / 2
		rightPad := padding - leftPad
		titleLine := "| " + strings.Repeat(" ", leftPad) + title + strings.Repeat(" ", rightPad) + " |\n"
		sb.WriteString(titleLine)
		sb.WriteString("+" + strings.Repeat("-", maxWidth+2) + "+\n")
	}

	for _, line := range lines {
		padding := maxWidth - len(line)
		sb.WriteString("| " + line + strings.Repeat(" ", padding) + " |\n")
	}

	sb.WriteString(bottomBorder)

	if ColorsEnabled {
		sb.WriteString(Reset)
	}

	return sb.String()
}

// Table creates a simple table
func Table(headers []string, rows [][]string) string {
	if len(headers) == 0 {
		return ""
	}

	// Calculate column widths
	colWidths := make([]int, len(headers))
	for i, h := range headers {
		colWidths[i] = len(h)
	}

	for _, row := range rows {
		for i, cell := range row {
			if i < len(colWidths) && len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	var sb strings.Builder

	// Header
	if ColorsEnabled {
		sb.WriteString(Bold)
	}
	for i, h := range headers {
		sb.WriteString(fmt.Sprintf("%-*s", colWidths[i]+2, h))
	}
	if ColorsEnabled {
		sb.WriteString(Reset)
	}
	sb.WriteString("\n")

	// Separator
	for _, w := range colWidths {
		sb.WriteString(strings.Repeat("-", w+2))
	}
	sb.WriteString("\n")

	// Rows
	for _, row := range rows {
		for i, cell := range row {
			if i < len(colWidths) {
				sb.WriteString(fmt.Sprintf("%-*s", colWidths[i]+2, cell))
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

// Spinner characters for animation
var spinnerChars = []string{"|", "/", "-", "\\"}

// Spinner returns a spinner character for the given frame
func Spinner(frame int) string {
	if !ColorsEnabled {
		return spinnerChars[frame%len(spinnerChars)]
	}
	return Cyan + spinnerChars[frame%len(spinnerChars)] + Reset
}

// Banner returns the application banner
func Banner(version string) string {
	banner := `
  __  __ _   _ ____  ____        ____  
 |  \/  | | | |  _ \|  _ \  ___ / ___|  
 | |\/| | |_| | | | | | | |/ _ \\___ \  
 | |  | |  _  | |_| | |_| | (_) |___) | 
 |_|  |_|_| |_|____/|____/ \___/|____/  
`
	if ColorsEnabled {
		return BrightCyan + banner + Reset + "\n" +
			Dim + "  DDoS Attack Script - Go Version " + version + Reset + "\n"
	}
	return banner + "\n  DDoS Attack Script - Go Version " + version + "\n"
}

// FormatMethod formats a method name with appropriate color
func FormatMethod(method string) string {
	return Color(BrightYellow, method)
}

// FormatTarget formats a target with appropriate color
func FormatTarget(target string) string {
	return Color(BrightMagenta, target)
}

// FormatNumber formats a number with appropriate color
func FormatNumber(num string) string {
	return Color(BrightGreen, num)
}

// ClearLine clears the current line in terminal
func ClearLine() {
	fmt.Print("\r\033[K")
}
