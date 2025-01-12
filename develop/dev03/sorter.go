package main

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

// Sorting options.
type sortOptions struct {
	column       int
	numeric      bool
	reverse      bool
	unique       bool
	month        bool
	ignoreSpaces bool
	check        bool
	humanNumeric bool
}

// Sorting strings.
func sortLines(lines []string, options sortOptions) ([]string, error) {
	if options.ignoreSpaces {
		for i, line := range lines {
			lines[i] = strings.TrimSpace(line)
		}
	}

	sort.SliceStable(lines, func(i, j int) bool {
		ki, kj := lines[i], lines[j]
		if options.column > 0 {
			ki = getColumn(ki, options.column)
			kj = getColumn(kj, options.column)
		}

		if options.numeric {
			return compareNumeric(ki, kj)
		} else if options.month {
			return compareMonth(ki, kj)
		} else if options.humanNumeric {
			return compareHumanNumeric(ki, kj)
		}
		return ki < kj
	})

	// if options.column > 0 {
	// 	sort.SliceStable(lines, func(i, j int) bool {
	// 		ki := getColumn(lines[i], options.column)
	// 		kj := getColumn(lines[j], options.column)
	//
	// 		if options.numeric {
	// 			return compareNumeric(ki, kj)
	// 		} else if options.month {
	// 			return compareMonth(ki, kj)
	// 		} else if options.humanNumeric {
	// 			return compareHumanNumeric(ki, kj)
	// 		}
	// 		return ki < kj
	// 	})
	// } else {
	// 	sort.SliceStable(lines, func(i, j int) bool {
	// 		if options.numeric {
	// 			return compareNumeric(lines[i], lines[j])
	// 		} else if options.month {
	// 			return compareMonth(lines[i], lines[j])
	// 		} else if options.humanNumeric {
	// 			return compareHumanNumeric(lines[i], lines[j])
	// 		}
	// 		return lines[i] < lines[j]
	// 	})
	// }

	if options.reverse {
		reverseSlice(lines)
	}

	if options.unique {
		lines = uniqueLines(lines)
	}

	return lines, nil
}

// compareNumeric - compare strings as numbers.
func compareNumeric(a, b string) bool {
	na, _ := strconv.ParseFloat(a, 64)
	nb, _ := strconv.ParseFloat(b, 64)
	return na < nb
}

// compareMonth - compare strings as months.
func compareMonth(a, b string) bool {
	ma, _ := time.Parse("Jan", a)
	mb, _ := time.Parse("Jan", b)
	return ma.Before(mb)
}

// compareHumanNumeric - comparison of "human-readable" numbers.
func compareHumanNumeric(a, b string) bool {
	na := parseHumanReadable(a)
	nb := parseHumanReadable(b)
	return na < nb
}

// parseHumanReadable - converts numbers with suffixes K, M, G, etc., into a numeric value.
func parseHumanReadable(s string) float64 {
	// Remove spaces.
	s = strings.TrimSpace(s)
	multiplier := 1.0

	// We determine the multiplier based on the last symbol.
	if len(s) > 1 {
		switch suffix := s[len(s)-1]; suffix {
		case 'K', 'k':
			multiplier = 1e3
			s = s[:len(s)-1]
		case 'M', 'm':
			multiplier = 1e6
			s = s[:len(s)-1]
		case 'G', 'g':
			multiplier = 1e9
			s = s[:len(s)-1]
		case 'T', 't':
			multiplier = 1e12
			s = s[:len(s)-1]
		}
	}

	// Parse the number and apply the multiplier.
	value, _ := strconv.ParseFloat(s, 64)
	return value * multiplier
}

// Check for sorting.
func isSorted(lines []string, options sortOptions) bool {
	for i := 1; i < len(lines); i++ {
		if options.reverse {
			if lines[i] > lines[i-1] {
				return false
			}
		} else {
			if lines[i] < lines[i-1] {
				return false
			}
		}
	}
	return true
}
