package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/icco/tcpfuzz"
)

func main() {
	port := "2222"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	log.Printf("Starting up on http://localhost:%s", port)

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalln("can't open:", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("can't accept: %+v", err)
			continue
		}

		go func(c net.Conn) {
			defer c.Close()
			log.Println("connection from", c.RemoteAddr())
			c.SetDeadline(time.Now().Add(10 * time.Second))
			err := tcpfuzz.WriteRand(c)
			if err != nil {
				log.Print(err)
			}
			return
		}(conn)
	}
}
