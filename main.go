package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
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
			WriteRand(c)
			return
		}(conn)
	}
}

func WriteRand(c net.Conn) {
	b := make([]byte, 1024*1024)
	i, err := rand.Read(b)
	if i < len(b) {
		log.Printf("could only get %d bytes of random", i)
	}

	if err != nil {
		log.Print(err)
	}

	c.Write(b)
}
