package tcpfuzz

import (
	"fmt"
	"math/rand"
	"net"
)

func WriteRand(c net.Conn) error {
	b := make([]byte, 1024*1024)
	i, err := rand.Read(b)
	if i < len(b) {
		return fmt.Errorf("could only get %d bytes of random", i)
	}

	if err != nil {
		return err
	}

	_, err = c.Write(b)
	return err
}
