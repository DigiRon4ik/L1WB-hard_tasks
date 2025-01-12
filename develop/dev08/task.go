package main

/*
=== Взаимодействие с ОС (ГАРАНТИРОВАННО РАБОТАЕТ ТОЛЬКО НА UNIX из-за os/exec) ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвейер на пайпах

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"sync"
)

var prompt = ""

func main() {
	fmt.Println("Simple UNIX Shell. Type \\quit to exit.")
	// Create an invitation.
	prompt = generatePrompt()
	reader := bufio.NewReader(os.Stdin)

	for {
		// Outputting an invitation.
		fmt.Print(prompt)
		// Read the command.
		input, err := reader.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}
		// Remove extra spaces and newlines.
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		// Processing the exit command.
		if input == "\\quit" {
			fmt.Println("Exiting shell.")
			break
		}

		err = execution(input)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			continue
		}
	}
}

// Generates invitation text.
func generatePrompt() string {
	curUser, err := user.Current()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error getting user:", err)
		os.Exit(1)
	}
	username := curUser.Username[strings.LastIndex(curUser.Username, "\\")+1:]
	hostname, err := os.Hostname()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error getting hostname:", err)
		os.Exit(1)
	}
	dir, err := os.Getwd()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error getting absolute path:", err)
	}
	lastSlash := strings.LastIndex(dir, "/")
	if lastSlash == -1 {
		lastSlash = strings.LastIndex(dir, "\\")
	}
	dir = dir[lastSlash:]

	return fmt.Sprintf("\n%s\n%s@%s ~ $ ", dir, username, hostname)
}

// Executing commands.
func execution(input string) error {
	commands := strings.Split(input, "|")
	cmds := make([]*exec.Cmd, len(commands))

	// Processing commands. If "cd", execute only it.
	for k, command := range commands {
		command = strings.TrimSpace(command)
		args := strings.Fields(command)
		if args[0] == "cd" {
			return changeDirectory(args)
		}
		cmd := exec.Command(args[0], args[1:]...)
		cmds[k] = cmd
	}

	var i int
	var wg sync.WaitGroup
	var err error

	// Set up a pipelines if there is one. And launch goroutines that execute commands.
	for i = 0; i < len(cmds)-1; i++ {
		r, w := io.Pipe()
		cmds[i].Stdout = w
		cmds[i].Stderr = os.Stderr
		cmds[i+1].Stdin = r
		wg.Add(1)
		go func(i int, w *io.PipeWriter) {
			defer func(w *io.PipeWriter) {
				_ = w.Close()
			}(w)
			err = cmds[i].Run()
			wg.Done()
			if err != nil {
				fmt.Println(err)
			}
		}(i, w)
	}
	// Configure the last (or only) command and launch the goroutine that executes the command.
	cmds[i].Stdout = os.Stdout
	cmds[i].Stderr = os.Stderr
	wg.Add(1)
	go func(i int) {
		err = cmds[i].Run()
		wg.Done()
		if err != nil {
			fmt.Println(err)
		}
	}(i)

	// Waiting for all commands to be completed.
	wg.Wait()
	return err
}

// cd command.
func changeDirectory(args []string) error {
	var err error

	if len(args) < 2 {
		// If the argument is not specified, change the directory to home.
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("could not get home directory: %v", err)
		}
		err = os.Chdir(homeDir)
	} else {
		// Change the directory to the specified one.
		err = os.Chdir(args[1])
	}
	prompt = generatePrompt()
	return err
}

/*
 - Usage (UNIX): -
go run task.go

 - Output: -
Simple UNIX Shell. Type \quit to exit.
/dev08
lux@KOMPUTER ~ $

 - Input: -
ls | sort -r | grep go

 - Output: -
task.go
go.mod
*/
