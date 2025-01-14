package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Parsing command line arguments.
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout (e.g., 5s, 10s, 1m)")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: go-telnet [--timeout=5s] host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	address := net.JoinHostPort(host, port)

	// Setting a timeout.
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Printf("Error connecting to %s: %v\n", address, err)
		os.Exit(1)
	}
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	fmt.Printf("Connected to %s\n", address)

	// Channel for completing work.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine for reading data from the connection and writing to STDOUT.
	go func() {
		buf := bufio.NewReader(conn)
		for {
			str, err := buf.ReadString('\n')
			if err != nil {
				fmt.Println("\nConnection closed by server.go.")
				os.Exit(0)
			}
			fmt.Print(str)
		}
	}()

	// Reading data from STDIN and writing to the connection.
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			str := scanner.Text()
			_, err := fmt.Fprintln(conn, str)
			if err != nil {
				fmt.Printf("Error writing to connection: %v\n", err)
				break
			} else if str == "exit" {
				break
			}
		}

		// Closing a socket by pressing Ctrl+D.
		if scanner.Err() == nil {
			fmt.Println("\nClosing connection...")
		}
		defer func(conn net.Conn) {
			_ = conn.Close()
		}(conn)
	}()

	// Waiting for completion.
	<-signals
	fmt.Printf("\nBye!\n")
}

/*
 - Usage: -
[UNIX]
go build -o go-telnet
go-telnet --timeout=2s nesuschestvujuschijSajt.ru 8080

[Windows]
go build -o go-telnet.exe
.\go-telnet.exe --timeout=2s nesuschestvujuschijSajt.ru 8080

[OR]
go run task.go --timeout=2s nesuschestvujuschijSajt.ru 8080

 - Output: -
Error connecting to nesuschestvujuschijSajt.ru:8080: dial tcp: lookup nesuschestvujuschijSajt.ru: i/o timeout
exit status 1
*/
