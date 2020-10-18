package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s [url] [...name]", os.Args[0])
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
	nodes := ElementsByTagName(doc, os.Args[2:]...)
	for _, node := range nodes {
		attrs := ""
		for _, attr := range node.Attr {
			attrs += fmt.Sprintf(" %s=\"%s\"", attr.Key, html.EscapeString(attr.Val))
		}
		fmt.Printf("<%s%s>\n", node.Data, attrs)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	return forEachNode(
		doc,
		name,
		func(n *html.Node, names []string) *html.Node {
			// fmt.Println(n.Data, names)
			if n.Type == html.ElementNode {
				for _, name := range names {
					if n.Data == name {
						return n
					}
				}
			}
			return nil
		},
		nil,
	)
}

func forEachNode(node *html.Node, names []string, pre, post func(n *html.Node, names []string) *html.Node) []*html.Node {
	nodes := make([]*html.Node, 0)
	if pre != nil {
		node := pre(node, names)
		if node != nil {
			nodes = append(nodes, node)
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, forEachNode(c, names, pre, post)...)
	}

	if post != nil {
		node := post(node, names)
		if node != nil {
			nodes = append(nodes, node)
		}
	}

	return nodes
}
