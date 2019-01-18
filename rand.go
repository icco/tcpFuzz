package tcpfuzz

import (
	"fmt"
	"io"
	"math/rand"
)

// WriteRand writes 1024^2 bytes into the provided io.Writer.
func WriteRand(c io.Writer) error {
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
