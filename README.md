# tcpFuzz

[![Go Report Card](https://goreportcard.com/badge/github.com/icco/tcpfuzz)](https://goreportcard.com/report/github.com/icco/tcpfuzz)

Just emits random content on a port. The idea is you create a port connection and then pass it in as an io.Writer into the random generator, and it will emit 1024kb of random data.

There are two example server directories above

 - ssh
 - tcp
