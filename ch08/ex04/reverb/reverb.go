package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer wg.Done()
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}

	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		time.Sleep(1 * time.Second)
		wg.Wait()
		os.Exit(0)
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		wg.Add(1)
		go handleConn(conn)
	}
}
