package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Parsing command line flags.
	column := flag.Int("k", 0, "column for sorting (1-based index)")
	numeric := flag.Bool("n", false, "sort numerically")
	reverse := flag.Bool("r", false, "reverse sort order")
	unique := flag.Bool("u", false, "output unique lines only")
	month := flag.Bool("M", false, "sort by month name")
	ignoreSpaces := flag.Bool("b", false, "ignore trailing spaces")
	check := flag.Bool("c", false, "check if file is sorted")
	humanNumeric := flag.Bool("h", false, "sort numerically with suffixes")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: sort [options] <input-file> [output-file]")
		return
	}

	inputFile := flag.Arg(0)
	outputFile := ""
	if flag.NArg() > 1 {
		outputFile = flag.Arg(1)
	}

	// Reading the input data.
	lines, err := readLines(inputFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		return
	}

	options := sortOptions{
		column:       *column,
		numeric:      *numeric,
		reverse:      *reverse,
		unique:       *unique,
		month:        *month,
		ignoreSpaces: *ignoreSpaces,
		check:        *check,
		humanNumeric: *humanNumeric,
	}

	// Checking the sorting if the -c flag is set.
	if options.check {
		if isSorted(lines, options) {
			fmt.Println("The file is sorted.")
		} else {
			fmt.Println("The file is not sorted.")
		}
		return
	}

	// Sort the lines.
	sortedLines, err := sortLines(lines, options)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error sorting lines: %v\n", err)
		return
	}

	// We write down the result.
	if outputFile != "" {
		err = writeLines(outputFile, sortedLines)
	} else {
		err = writeLinesToStdout(sortedLines)
	}
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
	}
}

/*
 - Usage: -
..\dev03-> go build
.\dev03.exe -k 2 -h -u -b input.txt output.txt

 - Input.txt: -
 Feb 10k
       Mar 5m
Jan 15
  Mar 5m
Jan 20
    Mar 5

 - Output.txt: -
Mar 5
Jan 15
Jan 20
Feb 10k
Mar 5m
*/
