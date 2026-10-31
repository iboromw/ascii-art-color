package main

import (
	"os"
	"strings"
)

// RenderColoredLines generates ASCII art for multiple lines, applying color to the specified substring.
func RenderColoredLines(lines []string, banner Banner, substring string, colorCode string) string {
	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		coloredIndices := getColoredIndices(line, substring)
		// If no substring is specified, color the entire line
		if substring == "" {
			for i := range line {
				coloredIndices[i] = true
			}
		}

		// Build the 8 lines of ASCII art for this text line
		for h := 0; h < 8; h++ {
			var lineBuilder strings.Builder
			inColor := false

			for i, ch := range line {
				charLines, exists := banner[ch]
				var charSlice string
				if exists {
					charSlice = charLines[h]
				} else {
					charSlice = banner[' '][h]
				}

				shouldColor := coloredIndices[i]

				if shouldColor && !inColor {
					lineBuilder.WriteString(colorCode)
					inColor = true
				} else if !shouldColor && inColor {
					lineBuilder.WriteString(ColorReset)
					inColor = false
				}

				lineBuilder.WriteString(charSlice)
			}

			if inColor {
				lineBuilder.WriteString(ColorReset)
			}

			result.WriteString(lineBuilder.String() + "\n")
		}
	}

	return result.String()
}

func getColoredIndices(line, substring string) []bool {
	colored := make([]bool, len(line))
	if substring == "" {
		return colored
	}

	start := 0
	for {
		idx := strings.Index(line[start:], substring)
		if idx == -1 {
			break
		}
		matchStart := start + idx
		matchEnd := matchStart + len(substring)
		for k := matchStart; k < matchEnd; k++ {
			colored[k] = true
		}
		start = matchStart + 1
	}
	return colored
}

func isBannerName(name string) bool {
	lower := strings.ToLower(name)
	if lower == "standard" || lower == "shadow" || lower == "thinkertoy" {
		return true
	}
	if _, err := os.Stat(ResolveBannerPath(name)); err == nil {
		return true
	}
	return false
}
