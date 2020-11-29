package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	message chan string
	ready   chan struct{}
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.message <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.message)
			close(cli.ready)
		}
	}
}

func handleConn(conn net.Conn) {
	cli := client{
		message: make(chan string),
		ready:   make(chan struct{}),
	}
	go clientWriter(conn, cli)

	who := conn.RemoteAddr().String()
	cli.message <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	cli.ready <- struct{}{}

loop:
	for {
		timer := time.NewTimer(5 * time.Second)
		c := make(chan string)
		go func() {
			if input.Scan() {
				c <- input.Text()
			} else {
				close(c)
			}
		}()
		select {
		case text, ok := <-c:
			timer.Stop()
			if !ok {
				break loop
			}
			messages <- who + ": " + text

		case <-timer.C:
			conn.Close()
			break loop
		}
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, cli client) {
	buffer := make([]string, 0)
standby:
	for {
		select {
		case _, ok := <-cli.ready:
			if !ok {
				return
			}
			break standby

		case msg, ok := <-cli.message:
			if !ok {
				return
			}
			buffer = append(buffer, msg)
		}
	}
	for _, msg := range buffer {
		fmt.Fprintln(conn, msg)
	}
	for msg := range cli.message {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
