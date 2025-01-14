package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// A simple server purely to test go-telnet.
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Can't connect !")
			_ = conn.Close()
			continue
		}
		remoteAddr := conn.RemoteAddr().String()
		fmt.Printf("%s Connected.\n", remoteAddr)
		fmt.Printf("%s : Start\n", remoteAddr)

		fmt.Printf("%s : Writing to console...\n", remoteAddr)
		_, err = conn.Write([]byte("\nHello from server!\n\n"))
		if err != nil {
			fmt.Println("Can't write!", err)
		}

		bufReader := bufio.NewReader(conn)
		go func(conn net.Conn) {
			defer func(conn net.Conn) {
				_ = conn.Close()
			}(conn)
			for {
				str, err := bufReader.ReadString('\n')
				if err != nil {
					fmt.Println("Can't read!", err)
					break
				}
				help := "\tshutdown\t- shut down the server\n" +
					"\tping\t\t- will print that the server is alive\n" +
					"\texit\t\t- logout from the server\n" +
					"\thelp\t\t- show this help\n"
				switch str {
				case "shutdown\n":
					fmt.Printf("%s : %s\n", remoteAddr, str)
					os.Exit(0)
				case "ping\n":
					_, _ = conn.Write([]byte("pong\n"))
				case "exit\n":
					_, _ = conn.Write([]byte("Come back soon!\n"))
				case "help\n":
					_, _ = conn.Write([]byte(help))
				default:
					_, _ = conn.Write([]byte("I don't understand. Type <help>\n"))
				}
				fmt.Printf("%s : %s", remoteAddr, str)
			}
		}(conn)
	}
}
