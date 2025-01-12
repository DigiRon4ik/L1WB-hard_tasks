package main

import (
	"bytes"
	"testing"
)

func TestProcessInput(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		delimiter    string
		separated    bool
		fieldIndexes []int
		expected     string
	}{
		{
			name:         "Single column with TAB delimiter",
			input:        "a\tb\tc\n1\t2\t3\n",
			delimiter:    "\t",
			separated:    false,
			fieldIndexes: []int{0},
			expected:     "a\n1\n",
		},
		{
			name:         "Multiple columns with TAB delimiter",
			input:        "a\tb\tc\n1\t2\t3\n",
			delimiter:    "\t",
			separated:    false,
			fieldIndexes: []int{0, 2},
			expected:     "a\tc\n1\t3\n",
		},
		{
			name:         "Custom delimiter",
			input:        "a,b,c\n1,2,3\n",
			delimiter:    ",",
			separated:    false,
			fieldIndexes: []int{1},
			expected:     "b\n2\n",
		},
		{
			name:         "Separated flag skips lines without delimiter",
			input:        "a:b:c\nno_delimiter\n1:2:3\n",
			delimiter:    ":",
			separated:    true,
			fieldIndexes: []int{0, 1},
			expected:     "a:b\n1:2\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := bytes.NewBufferString(tt.input)
			output := &bytes.Buffer{}

			err := processInput(input, output, tt.delimiter, tt.separated, tt.fieldIndexes)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Fatalf("expected: %q, got: %q", tt.expected, output.String())
			}
		})
	}
}

func TestParseFields(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []int
		expectedErr bool
	}{
		{"Valid single field", "1", []int{0}, false},
		{"Valid multiple fields", "1,3,5", []int{0, 2, 4}, false},
		{"Invalid field format", "a,b,c", nil, true},
		{"Negative field number", "-1", nil, true},
		{"Zero field number", "0", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseFields(tt.input)
			if (err != nil) != tt.expectedErr {
				t.Fatalf("expected error: %v, got: %v", tt.expectedErr, err)
			}
			if !equalIntSlices(result, tt.expected) {
				t.Fatalf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

// equalIntSlices compares two integer slices.
func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
