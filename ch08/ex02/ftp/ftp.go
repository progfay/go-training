package ftp

import (
	"bufio"
	"fmt"
	"log"
	"net"
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
	log.Println(input)
	for input.Scan() {
		fmt.Println(input.Text())
	}
}
