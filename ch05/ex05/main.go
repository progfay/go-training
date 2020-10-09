package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [url]\n", os.Args[0])
		os.Exit(1)
	}

	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("url\t%s\n", os.Args[1])
	fmt.Printf("words\t%d\n", words)
	fmt.Printf("images\t%d\n", images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}

	switch n.Type {
	case html.ElementNode:
		if n.Data == "img" {
			images++
		}
		w, i := countWordsAndImages(n.NextSibling)
		words, images = words+w, images+i
		w, i = countWordsAndImages(n.FirstChild)
		words, images = words+w, images+i

	case html.TextNode:
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}

	return
}
