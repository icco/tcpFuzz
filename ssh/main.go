package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gliderlabs/ssh"
	"github.com/icco/tcpfuzz"
)

func main() {
	port := "2222"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	log.Printf("starting ssh server on port %s...", port)
	log.Fatal(ssh.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
		err := tcpfuzz.WriteRand(s)
		if err != nil {
			log.Print(err)
		}
	}))
}
