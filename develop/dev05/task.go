package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Parsing command line flags.
	after := flag.Int("A", 0, "Print +N lines after match")
	before := flag.Int("B", 0, "Print +N lines before match")
	context := flag.Int("C", 0, "Print ±N lines around match (overrides -A and -B)")
	count := flag.Bool("c", false, "Print count of matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert match")
	fixed := flag.Bool("F", false, "Match fixed string (not pattern)")
	lineNum := flag.Bool("n", false, "Print line number")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: grep [OPTIONS] PATTERN [FILE]")
		os.Exit(1)
	}

	pattern := flag.Arg(0)
	var inputFile string

	// Reading a file or stdin.
	if flag.NArg() > 1 {
		inputFile = flag.Arg(1)
	}

	lines, err := ReadLinesOrStdin(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Applying filtering.
	config := Config{
		After:      *after,
		Before:     *before,
		Context:    *context,
		Count:      *count,
		IgnoreCase: *ignoreCase,
		Invert:     *invert,
		Fixed:      *fixed,
		LineNum:    *lineNum,
		Pattern:    pattern,
	}
	result := ProcessLines(lines, config)

	// Output of results.
	for _, line := range result {
		fmt.Println(line)
	}
}

/*
 - Usage: -
..\dev05-> go build
.\dev05.exe -i -n go input.txt

 - Input.txt: -
Go is great.
Python is good.
I love programming.
Go is fast.
Go is simple.

 - Output: -
1: Go is great.
2: Python is good.
4: Go is fast.
5: Go is simple.
*/
