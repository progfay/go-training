package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(c net.Conn, timezone string) {
	defer c.Close()
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc = time.FixedZone(timezone, 0)
	}
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	port := flag.Int("port", 8000, "listen port")
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		timezone := os.Getenv("TZ")
		go handleConn(conn, timezone)
	}
}
