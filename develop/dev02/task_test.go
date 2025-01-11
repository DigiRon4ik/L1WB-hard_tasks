package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{`qwe\4\5`, "qwe45", false},
		{`qwe\45`, "qwe44444", false},
		{`qwe\\5`, `qwe\\\\\`, false},
		{`\`, "", true},  // Invalid string: incomplete escape sequence.
		{`a\`, "", true}, // Invalid string: incomplete escape sequence.
		{`a4b\2`, "aaaab2", false},
		{`a4\`, "", true},       // Invalid string: incomplete escape sequence.
		{"123abc", "", true},    // Invalid string: starts with a number.
		{`\\3a`, `\\\a`, false}, // Escape symbol with repetition.
	}

	for _, test := range tests {
		result, err := UnpackString(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("UnpackString(%q) unexpected error status: got %v, want error: %v", test.input, err, test.hasError)
		}
		if result != test.expected {
			t.Errorf("UnpackString(%q) = %q, want %q", test.input, result, test.expected)
		}
	}
}
