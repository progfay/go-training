package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// Ref. https://www.w3.org/TR/REC-html40/index/attributes.html
var urlTypeHtmlAttributes = []string{
	"action",
	"background",
	"cite",
	"classid",
	"codebase",
	"data",
	"href",
	"longdesc",
	"profile",
	"src",
	"usemap",
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	links := Visit(nil, doc)
	for _, text := range links {
		fmt.Printf("%#v\n", text)
	}
}

func Visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			for _, urlTypeHtmlAttribute := range urlTypeHtmlAttributes {
				if attr.Key == urlTypeHtmlAttribute {
					links = append(links, attr.Val)
				}
			}
		}
	}

	links = Visit(links, n.NextSibling)
	links = Visit(links, n.FirstChild)

	return links
}
