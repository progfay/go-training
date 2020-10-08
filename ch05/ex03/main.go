package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	texts := Visit(nil, doc)
	for _, text := range texts {
		fmt.Printf("%#v\n", text)
	}
}

func Visit(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}

	if n.Type == html.TextNode {
		data := strings.TrimSpace(n.Data)
		if data != "" {
			texts = append(texts, data)
		}
	}

	texts = Visit(texts, n.NextSibling)

	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return texts
		}
	}

	return Visit(texts, n.FirstChild)
}
