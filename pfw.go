package main

import (
	"io"
	"log"
	"net"
	"os"
)

func forward(conn net.Conn) {
	client, err := net.Dial("tcp", os.Args[2])
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	log.Printf("Connected to localhost %v\n", conn)
	go func() {
		defer client.Close()
		defer conn.Close()
		io.Copy(client, conn)
	}()
	go func() {
		defer client.Close()
		defer conn.Close()
		io.Copy(conn, client)
	}()
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage %s listen:port forward:port\n", os.Args[0]);
		return
	}

	listener, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatalf("Failed to setup listener: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("ERROR: failed to accept listener: %v", err)
		}
		log.Printf("Accepted connection %v\n", conn)
		go forward(conn)
	}
}
