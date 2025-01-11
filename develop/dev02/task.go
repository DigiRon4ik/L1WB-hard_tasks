package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// UnpackString unpacks a string according to the specified rules.
func UnpackString(input string) (string, error) {
	var (
		result   strings.Builder
		prevRune rune
		escaping bool
	)

	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		current := runes[i]

		switch {
		case escaping:
			// If the previous character was '\', then add the current one as is.
			result.WriteRune(current)
			prevRune = current
			escaping = false

		case current == '\\':
			// Process the beginning of the escape sequence.
			escaping = true

		case unicode.IsDigit(current):
			if prevRune == 0 {
				return "", errors.New("input string is invalid: cannot start with a digit")
			}

			// Count multi-digit numbers.
			digitStart := i
			for i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
				i++
			}
			number, err := strconv.Atoi(string(runes[digitStart : i+1]))
			if err != nil {
				return "", fmt.Errorf("failed to parse repeat count: %w", err)
			}

			// Add repetitions of the previous rune.
			result.WriteString(strings.Repeat(string(prevRune), number-1))

		default:
			// If the character is not a digit, neither escape nor repeat.
			result.WriteRune(current)
			prevRune = current
		}
	}

	// Check for incomplete escape.
	if escaping {
		return "", errors.New("input string is invalid: incomplete escape sequence")
	}

	return result.String(), nil
}

func main() {
	fmt.Println(UnpackString(`a2bc10d1e\3\\10zxc\25`))
}

/*
 - Output: -
aabccccccccccde3\\\\\\\\\\zxc22222 <nil>
*/
