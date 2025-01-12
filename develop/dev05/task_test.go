package main

import (
	"testing"
)

func TestProcessLines(t *testing.T) {
	lines := []string{
		"Go is great.",
		"Python is good.",
		"I love programming.",
		"Go is fast.",
		"Go is simple.",
	}

	tests := []struct {
		name   string
		config Config
		expect []string
	}{
		{
			name:   "Basic match",
			config: Config{Pattern: "Go"},
			expect: []string{"Go is great.", "Go is fast.", "Go is simple."},
		},
		{
			name:   "Case insensitive match",
			config: Config{Pattern: "go", IgnoreCase: true},
			expect: []string{"Go is great.", "Python is good.", "Go is fast.", "Go is simple."},
		},
		{
			name:   "Invert match",
			config: Config{Pattern: "Go", Invert: true},
			expect: []string{"Python is good.", "I love programming."},
		},
		{
			name:   "Context match",
			config: Config{Pattern: "Go", Context: 1},
			expect: []string{"Go is great.", "Python is good.", "I love programming.", "Go is fast.", "Go is simple."},
		},
		{
			name:   "Count match",
			config: Config{Pattern: "Go", Count: true},
			expect: []string{"3"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ProcessLines(lines, test.config)
			if len(result) != len(test.expect) {
				t.Errorf("Expected %d lines, got %d", len(test.expect), len(result))
			}
			for i, line := range result {
				if line != test.expect[i] {
					t.Errorf("Expected '%s', got '%s'", test.expect[i], line)
				}
			}
		})
	}
}
