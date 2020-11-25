package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/progfay/go-training/ch08/ex06/links"
)

type Link struct {
	depth int64
	url   string
}

func crawl(link Link) []string {
	fmt.Println(link.url)
	list, err := links.Extract(link.url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	depth := flag.Int64("depth", 3, "depth limit of crawling")
	flag.Parse()

	worklist := make(chan []Link)
	unseenLinks := make(chan Link)

	seed := make([]Link, len(os.Args)-1)
	for i, arg := range os.Args[1:] {
		seed[i] = Link{
			depth: 0,
			url:   arg,
		}
	}

	go func() { worklist <- seed }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				if link.depth >= *depth {
					return
				}
				foundUrls := crawl(link)
				linklist := make([]Link, 0)
				for _, url := range foundUrls {
					linklist = append(linklist, Link{
						depth: link.depth + 1,
						url:   url,
					})
				}
				go func() { worklist <- linklist }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}
