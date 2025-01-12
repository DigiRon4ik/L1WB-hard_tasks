package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Parsing command line arguments.
	fields := flag.String("f", "", "select fields (columns)")
	delimiter := flag.String("d", "\t", "use a different delimiter (default is TAB)")
	separated := flag.Bool("s", false, "only delimited lines")
	flag.Parse()

	// Checks that the -f argument is not empty.
	if *fields == "" {
		_, _ = fmt.Fprintln(os.Stderr, "Error: -f key is required")
		flag.Usage()
		os.Exit(1)
	}

	// Parse the list of columns.
	fieldIndexes, err := parseFields(*fields)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: Invalid argument format -f: %v\n", err)
		os.Exit(1)
	}

	// Process STDIN and output the result to STDOUT.
	if err := processInput(os.Stdin, os.Stdout, *delimiter, *separated, fieldIndexes); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Data processing error: %v\n", err)
		os.Exit(1)
	}
}

// processInput performs basic processing of input data.
func processInput(input io.Reader, output io.Writer, delimiter string, separated bool, fieldIndexes []int) error {
	scanner := bufio.NewScanner(input)
	writer := bufio.NewWriter(output)
	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)

	for scanner.Scan() {
		line := scanner.Text()

		// Check for the presence of a delimiter if the -s flag is set.
		if separated && !strings.Contains(line, delimiter) {
			continue
		}

		// Split the string into parts.
		columns := strings.Split(line, delimiter)

		// Select only the specified columns.
		selectedColumns := selectFields(columns, fieldIndexes)

		// Write down the result.
		_, err := writer.WriteString(strings.Join(selectedColumns, delimiter) + "\n")
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}

// parseFields parses a string with column numbers into an array of indexes.
func parseFields(fields string) ([]int, error) {
	parts := strings.Split(fields, ",")
	indexes := make([]int, 0, 8)
	for _, part := range parts {
		index, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		if index < 1 {
			return nil, fmt.Errorf("column number must be > 0")
		}
		indexes = append(indexes, index-1)
	}
	return indexes, nil
}

// selectFields selects from the list of columns only those specified in fieldIndexes.
func selectFields(columns []string, fieldIndexes []int) []string {
	result := make([]string, 0, 8)
	for _, index := range fieldIndexes {
		if index >= 0 && index < len(columns) {
			result = append(result, columns[index])
		}
	}
	return result
}

/*
 - Usage: -
[pwsh]
..\dev03-> "a:b:c`nno_delimiter`n1:2:3" | go run task.go -f 1,2 -d ":" -s
[bash]
../dev03$ echo -e "a:b:c\nno_delimiter\n1:2:3" | go run task.go -f 1,2 -d ":" -s

 - Output: -
a:b
1:2
*/
