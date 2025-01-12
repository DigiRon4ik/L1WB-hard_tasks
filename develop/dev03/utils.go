package main

import (
	"bufio"
	"os"
	"strings"
)

// Reading lines from a file.
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0, 32)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Writing lines to a file.
func writeLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// Writing strings to stdout.
func writeLinesToStdout(lines []string) error {
	for _, line := range lines {
		println(line)
	}
	return nil
}

// Getting a specific column.
func getColumn(line string, column int) string {
	parts := strings.Fields(line)
	if column-1 < len(parts) {
		return parts[column-1]
	}
	return ""
}

// Reverse line.
func reverseSlice(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Unique strings.
func uniqueLines(lines []string) []string {
	seen := make(map[string]struct{})
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		if _, exists := seen[line]; !exists {
			seen[line] = struct{}{}
			result = append(result, line)
		}
	}
	return result
}
