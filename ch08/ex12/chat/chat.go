package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string
type user struct {
	client client
	who    string
}

var (
	entering = make(chan user)
	leaving  = make(chan user)
	messages = make(chan string)
)

var clients map[client]string

func broadcaster() {
	clients = make(map[client]string)
	for {
		select {
		case msg := <-messages:

			for cli := range clients {
				cli <- msg
			}

		case user := <-entering:
			clients[user.client] = user.who

		case user := <-leaving:
			delete(clients, user.client)
			close(user.client)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "============================"
	ch <- "You are " + who
	ch <- "members:"
	for _, name := range clients {
		ch <- "- " + name
	}
	ch <- "============================"
	messages <- who + " has arrived"
	entering <- user{ch, who}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- user{ch, who}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
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
