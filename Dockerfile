FROM golang:1.11
ENV GO111MODULE=on
ENV PORT 22
EXPOSE 22
WORKDIR /go/src/github.com/icco/tcpfuzz
COPY . .

RUN go build -o /go/bin/server ./ssh

CMD ["/go/bin/server"]
