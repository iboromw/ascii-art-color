package main

// GetColorCode maps color names or codes to ANSI escape sequences.
func GetColorCode(color string) string {
	switch color {
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "magenta", "purple":
		return "\033[35m"
	case "cyan":
		return "\033[36m"
	case "white":
		return "\033[37m"
	default:
		// Default to red if unrecognized
		return "\033[31m"
	}
}

const ColorReset = "\033[0m"
