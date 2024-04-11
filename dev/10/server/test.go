package main

import (
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 32)
	for {
		conn.Write([]byte("Hello, what's your name?\n"))

		readLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}

		conn.Write(append([]byte("Goodbye, "), buf[:readLen]...))
	}
}
