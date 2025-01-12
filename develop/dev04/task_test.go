package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		{
			name:  "Basic test with multiple groups",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "Слиток", "ПЯТАК"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:  "Case insensitive and duplicates",
			input: []string{"графин", "нифгра", "НифГра", "графин"},
			expected: map[string][]string{
				"графин": {"графин", "нифгра"},
			},
		},
		{
			name:     "No anagrams",
			input:    []string{"слово", "текст", "код"},
			expected: map[string][]string{},
		},
		{
			name:  "Mixed single and multiple anagram groups",
			input: []string{"один", "дино", "два", "адв", "три", "рит"},
			expected: map[string][]string{
				"один": {"дино", "один"},
				"два":  {"адв", "два"},
				"три":  {"рит", "три"},
			},
		},
		{
			name:     "Empty input",
			input:    []string{},
			expected: map[string][]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findAnagramSets(&test.input)
			if !reflect.DeepEqual(*result, test.expected) {
				t.Errorf("Unexpected result. Got: %v, Expected: %v", *result, test.expected)
			}
		})
	}
}

func TestSortString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"пятак", "акптя"},
		{"слово", "влоос"},
		{"графин", "агинрф"},
		{"Текст", "Текст"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := sortString(test.input)
			if result != test.expected {
				t.Errorf("Unexpected result for %s. Got: %s, Expected: %s", test.input, result, test.expected)
			}
		})
	}
}

func TestUniqueAndSort(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"графин", "нифгра", "графин"}, []string{"графин", "нифгра"}},
		{[]string{"слово", "текст", "слово"}, []string{"слово", "текст"}},
		{[]string{"один", "один", "один"}, []string{"один"}},
		{[]string{}, []string{}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := uniqueAndSort(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Unexpected result for %v. Got: %v, Expected: %v", test.input, result, test.expected)
			}
		})
	}
}
