package main

import (
	"testing"
)

func equalSlices(a, b []string) bool {
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

func TestSortLinesByNumericValue(t *testing.T) {
	lines := []string{"10", "2", "1", "5", "20"}
	expected := []string{"1", "2", "5", "10", "20"}

	options := sortOptions{numeric: true}
	result, _ := sortLines(lines, options)

	if !equalSlices(result, expected) {
		t.Errorf("Numeric sort failed. Expected %v, got %v", expected, result)
	}
}

func TestSortLinesReverse(t *testing.T) {
	lines := []string{"a", "b", "c"}
	expected := []string{"c", "b", "a"}

	options := sortOptions{reverse: true}
	result, _ := sortLines(lines, options)

	if !equalSlices(result, expected) {
		t.Errorf("Reverse sort failed. Expected %v, got %v", expected, result)
	}
}

func TestUniqueLines(t *testing.T) {
	lines := []string{"a", "b", "a", "c", "b"}
	expected := []string{"a", "b", "c"}

	options := sortOptions{unique: true}
	result, _ := sortLines(lines, options)

	if !equalSlices(result, expected) {
		t.Errorf("Unique sort failed. Expected %v, got %v", expected, result)
	}
}

func TestSortLinesByColumn(t *testing.T) {
	lines := []string{"c 3", "a 1", "b 2"}
	expected := []string{"a 1", "b 2", "c 3"}

	options := sortOptions{column: 2, numeric: true}
	result, _ := sortLines(lines, options)

	if !equalSlices(result, expected) {
		t.Errorf("Column sort failed. Expected %v, got %v", expected, result)
	}
}

func TestSortLinesByMonth(t *testing.T) {
	lines := []string{"Feb", "Jan", "Mar"}
	expected := []string{"Jan", "Feb", "Mar"}

	options := sortOptions{month: true}
	result, _ := sortLines(lines, options)

	if !equalSlices(result, expected) {
		t.Errorf("Month sort failed. Expected %v, got %v", expected, result)
	}
}

func TestTrimTrailingSpaces(t *testing.T) {
	lines := []string{"a ", "b  ", " c"}
	expected := []string{"a", "b", "c"}

	options := sortOptions{ignoreSpaces: true}
	result, _ := sortLines(lines, options)

	if !equalSlices(result, expected) {
		t.Errorf("Ignore trailing spaces failed. Expected %v, got %v", expected, result)
	}
}

func TestIsSorted(t *testing.T) {
	lines := []string{"a", "b", "c"}
	options := sortOptions{}

	if !isSorted(lines, options) {
		t.Errorf("IsSorted failed. Expected true, got false")
	}

	lines = []string{"c", "b", "a"}
	options = sortOptions{reverse: true}

	if !isSorted(lines, options) {
		t.Errorf("IsSorted failed. Expected true, got false")
	}
}

func TestSortLinesHumanReadable(t *testing.T) {
	lines := []string{"1K", "2M", "500", "3G", "10K"}
	expected := []string{"500", "1K", "10K", "2M", "3G"}

	options := sortOptions{humanNumeric: true}
	result, _ := sortLines(lines, options)

	if !equalSlices(result, expected) {
		t.Errorf("Human-readable numeric sort failed. Expected %v, got %v", expected, result)
	}
}

func TestParseHumanReadable(t *testing.T) {
	tests := map[string]float64{
		"1K":  1000,
		"2M":  2000000,
		"3G":  3000000000,
		"500": 500,
		"10K": 10000,
	}

	for input, expected := range tests {
		result := parseHumanReadable(input)
		if result != expected {
			t.Errorf("ParseHumanReadable failed. Input: %s, Expected: %v, Got: %v", input, expected, result)
		}
	}
}

func TestGetColumn(t *testing.T) {
	line := "a b c"
	expected := "b"

	result := getColumn(line, 2)
	if result != expected {
		t.Errorf("GetColumn failed. Expected %v, got %v", expected, result)
	}
}

func TestReverseSlice(t *testing.T) {
	lines := []string{"a", "b", "c"}
	expected := []string{"c", "b", "a"}

	reverseSlice(lines)
	if !equalSlices(lines, expected) {
		t.Errorf("ReverseSlice failed. Expected %v, got %v", expected, lines)
	}
}

func TestUniqueLinesFunction(t *testing.T) {
	lines := []string{"a", "b", "a", "c"}
	expected := []string{"a", "b", "c"}

	result := uniqueLines(lines)
	if !equalSlices(result, expected) {
		t.Errorf("UniqueLines failed. Expected %v, got %v", expected, result)
	}
}
