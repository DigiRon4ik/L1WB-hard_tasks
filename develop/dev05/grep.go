package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Config хранит параметры конфигурации для фильтрации
type Config struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
	Pattern    string
}

// ReadLinesOrStdin reading lines from a file or stdin.
func ReadLinesOrStdin(filename string) ([]string, error) {
	var scanner *bufio.Scanner

	if filename != "" {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	lines := make([]string, 0, 32)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// ProcessLines filters lines based on a given configuration.
func ProcessLines(lines []string, config Config) []string {
	if config.IgnoreCase {
		config.Pattern = strings.ToLower(config.Pattern)
	}

	var matches []int // Indexes of rows matching the filter.
	for i, line := range lines {
		var lineToCheck string
		if config.IgnoreCase {
			lineToCheck = strings.ToLower(line)
		} else {
			lineToCheck = line
		}

		match := false
		if config.Fixed {
			match = lineToCheck == config.Pattern
		} else {
			match = strings.Contains(lineToCheck, config.Pattern)
		}

		if config.Invert {
			match = !match
		}

		if match {
			matches = append(matches, i)
		}
	}

	// If the -c flag is specified, return the number of matches.
	if config.Count {
		return []string{fmt.Sprintf("%d", len(matches))}
	}

	// Form a conclusion taking into account the context.
	output := make(map[int]bool, len(matches))
	for _, match := range matches {
		start := match - config.Before
		if config.Context > 0 {
			start = match - config.Context
		}
		if start < 0 {
			start = 0
		}

		end := match + config.After
		if config.Context > 0 {
			end = match + config.Context
		}
		if end >= len(lines) {
			end = len(lines) - 1
		}

		for i := start; i <= end; i++ {
			output[i] = true
		}
	}

	// Generate the result.
	result := make([]string, 0, len(output))
	for i, line := range lines {
		if output[i] {
			if config.LineNum {
				result = append(result, fmt.Sprintf("%d: %s", i+1, line))
			} else {
				result = append(result, line)
			}
		}
	}

	return result
}
