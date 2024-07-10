package main

import (
	"io"
	"log"
	"net"
)

var (
	// DefaultHTTPGetAddress Default Address
	listenAddr = "localhost:8080"

	server = []string{
		"localhost:5050",
		"localhost:5051",
		"localhost:5052",
	}
	count int
)

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("cannot listen: %s", err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("WARNING: cannot accept connection err: %s", err)
		}
		backendServer := chooseServer()
		go func() {
			err := proxy(conn, backendServer)
			if err != nil {
				log.Printf("Failed to connect proxy server %s", err)
			}
		}()
	}

}

func proxy(c net.Conn, server string) error {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Printf("WARNING: proxy server is unavailable %s", err)
	}
	go io.Copy(conn, c)
	go io.Copy(c, conn)

	return nil
}

func chooseServer() string {
	s := server[count%len(server)]
	count++
	return s
}
