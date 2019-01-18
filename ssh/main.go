package main

import (
	"fmt"
	"io"
	"log"

	"github.com/gliderlabs/ssh"
	"github.com/icco/tcpfuzz"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
		err := tcpfuzz.WriteRand(s)
		if err != nil {
			log.Print(err)
		}
	})

	log.Println("starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
