package ftp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type ftpServer struct {
	listener net.Listener
	close    chan struct{}
}

func New(url string) (*ftpServer, error) {
	listener, err := net.Listen("tcp", url)
	if err != nil {
		return nil, err
	}

	return &ftpServer{
		listener: listener,
		close:    make(chan struct{}),
	}, nil
}

func (s *ftpServer) Listen() {
	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				log.Println(err)
				break
			}

			go handleConnection(conn)
		}
		close(s.close)
	}()
}

func (s *ftpServer) Close() {
	close(s.close)
}

func (s *ftpServer) Cancel() <-chan struct{} {
	return s.close
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	input := bufio.NewScanner(conn)
	fmt.Fprintln(conn, 220)

	for input.Scan() {
		req := parse(input.Text())
		fmt.Printf("%#v\n", req)
	}
}

type request struct {
	command string
	message string
}

func parse(text string) request {
	s := strings.SplitN(text, " ", 2)

	switch len(s) {
	case 0:
		return request{}

	case 1:
		return request{
			command: s[0],
		}

	default:
		return request{
			command: s[0],
			message: s[1],
		}
	}
}
