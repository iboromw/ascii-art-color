package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	// Validate minimum argument length
	if len(args) < 2 || len(args) > 4 {
		printUsage()
		return
	}

	flag := args[0]
	if !strings.HasPrefix(flag, "--color=") {
		printUsage()
		return
	}

	colorName := strings.TrimPrefix(flag, "--color=")
	colorCode := GetColorCode(colorName)

	var substring, rawInput, bannerName string
	bannerName = "standard" // default banner

	// Parse argument combinations intelligently
	if len(args) == 2 {
		rawInput = args[1]
	} else if len(args) == 3 {
		if isBannerName(args[2]) {
			rawInput = args[1]
			bannerName = args[2]
		} else {
			substring = args[1]
			rawInput = args[2]
		}
	} else if len(args) == 4 {
		substring = args[1]
		rawInput = args[2]
		bannerName = args[3]
	}

	if rawInput == "" {
		return
	}

	if err := ValidateInput(rawInput); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	banner, err := LoadBanner(bannerName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	lines := ParseInput(rawInput)
	if len(lines) == 0 {
		return
	}

	output := RenderColoredLines(lines, banner, substring, colorCode)
	fmt.Print(output)
}

func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING]  EX: go run . --color=<color> <substring to be colored> \"something\"")
}
