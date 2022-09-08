package main

import (
	"io"
	"log"
	"net"
	"os"
)

// This function copies the destination from the standard output and read the source from reader interface
func copy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// connect to server on localhost 80
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	copy(os.Stdout, conn)
}
