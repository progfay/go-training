package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"

	"github.com/progfay/go-training/ch05/ex07/prettier"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s [url] [id]", os.Args[0])
	}

	resp, err := http.Get(os.Args[1])
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
	node := ElementByID(doc, os.Args[2])
	prettier.WritePrettyHtml(os.Stdout, node)
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, endElement)
}

func forEachNode(node *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	var exit bool
	if pre != nil {
		exit = pre(node, id)
		if exit {
			return node
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		n := forEachNode(c, id, pre, post)
		if n != nil {
			return n
		}
	}

	if post != nil {
		exit = post(node, id)
		if exit {
			return node
		}
	}

	return nil
}

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return true
			}
		}
	}
	return false
}

func endElement(n *html.Node, id string) bool {
	return false
}
