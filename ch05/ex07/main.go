package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"

	"github.com/progfay/go-training/ch05/ex07/prettier"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [...url]", os.Args[0])
	}

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		prettier.WritePrettyHtml(os.Stdout, doc)
	}
}
