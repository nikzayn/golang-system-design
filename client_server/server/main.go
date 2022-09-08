package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "Hi from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}

func main() {
	// Handling concurrent client operations

	// Listening on port 8080 as default
	fmt.Println("Server start on port 80")
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Accepting connection on port 80
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}
