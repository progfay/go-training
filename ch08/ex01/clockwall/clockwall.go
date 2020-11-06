package main

import (
	"log"
	"os"
	"strings"
	"sync"
	"text/tabwriter"

	"github.com/progfay/go-training/ch08/ex01/clockwall/connection"
)

func main() {
	connections := make([]*connection.Connection, len(os.Args)-1)
	for i, arg := range os.Args[1:] {
		c, err := connection.New(arg)
		if err != nil {
			log.Fatal(err)
		}
		defer c.Close()
		connections[i] = c
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 8, 0, '\t', 0)

	for _, c := range connections {
		w.Write([]byte(c.Name + "\t"))
	}
	w.Write([]byte("\n\r"))
	w.Flush()

	times := make([]string, len(connections))
	mu := sync.Mutex{}
	errChan := make(chan error)

	for i, conn := range connections {
		go func(i int, conn *connection.Connection) {
			for {
				var buf = make([]byte, 8)
				_, err := conn.Read(buf)
				if err != nil {
					errChan <- err
				}
				t := strings.Trim(string(buf), "\n\x00")
				if len(t) == 0 {
					continue
				}

				func() {
					mu.Lock()
					defer mu.Unlock()
					times[i] = t
					_, err := w.Write([]byte(strings.Join(times, "\t") + "\r"))
					if err != nil {
						errChan <- err
					}
					err = w.Flush()
					if err != nil {
						errChan <- err
					}
				}()
			}
		}(i, conn)
	}

	select {
	case err := <-errChan:
		log.Fatal(err)
	}
}
