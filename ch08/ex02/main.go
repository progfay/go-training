package main

import (
	"fmt"
	"log"
	"os"

	"github.com/progfay/go-training/ch08/ex02/ftp"
)

func main() {
	fptServer, err := ftp.New("localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	fptServer.Listen()

	pressed := make(chan struct{})
	go func() {
		fmt.Println("to shutdown server, press ENTER key...")
		os.Stdin.Read(make([]byte, 1))
		pressed <- struct{}{}
	}()

	select {
	case <-pressed:
		fptServer.Close()

	case <-fptServer.Cancel():
	}
}
