package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	who := flag.String("who", "anonymous", "username")
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	c, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatalf("conn is not TCPConn: %v", conn)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, c)
		c.CloseRead()
		log.Println("done")
		done <- struct{}{}
	}()
	fmt.Fprintln(c, *who)
	mustCopy(c, os.Stdin)
	c.CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
