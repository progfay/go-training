package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	countMap := make(map[string]int)
	Visit(countMap, doc)
	for data, count := range countMap {
		fmt.Printf("%s\t%d\n", data, count)
	}
}

func Visit(countMap map[string]int, n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		countMap[n.Data]++
	}

	Visit(countMap, n.FirstChild)
	Visit(countMap, n.NextSibling)
}
