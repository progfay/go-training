package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/progfay/go-training/ch08/ex10/links"
)

func crawl(link string, exit <-chan struct{}, wg *sync.WaitGroup) []string {
	wg.Add(1)
	defer wg.Done()

	fmt.Println(link)
	list, err := links.Extract(link, exit)
	if err != nil {
		select {
		case <-exit:
			log.Printf("request was cancelled: %q\n", link)

		default:
			log.Println(err)
		}
	}
	return list
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	exit := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		log.Println("exit app, start cancelling ongoing requests...")
		close(exit)
	}()

	var wg sync.WaitGroup

	go func() {
		select {
		case <-exit:

		default:
			worklist <- os.Args[1:]
		}
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundUrls := crawl(link, exit, &wg)
				select {
				case <-exit:
					return

				default:
				}
				linklist := make([]string, 0)
				for _, url := range foundUrls {
					linklist = append(linklist, url)
				}
				go func() { worklist <- linklist }()
			}
		}()
	}

	seen := make(map[string]bool)

	for {
		select {
		case list := <-worklist:
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}

		case <-exit:
			wg.Wait()
			os.Exit(0)
		}
	}
}
