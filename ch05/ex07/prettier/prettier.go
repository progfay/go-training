package prettier

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func WritePrettyHtml(w io.Writer, doc *html.Node) {
	forEachNode(w, doc, 0, startElement, endElement)
}

func forEachNode(w io.Writer, n *html.Node, depth int, pre, post func(w io.Writer, n *html.Node, depth int) int) {
	if pre != nil {
		depth = pre(w, n, depth)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(w, c, depth, pre, post)
	}

	if post != nil {
		depth = post(w, n, depth)
	}
}

func startElement(w io.Writer, n *html.Node, depth int) int {
	switch n.Type {
	case html.TextNode, html.CommentNode:
		data := strings.TrimSpace(n.Data)
		if data != "" {
			fmt.Fprintf(w, "%*s%s\n", depth*2, "", data)
		}

	case html.ElementNode:
		attributes := ""
		for _, attr := range n.Attr {
			attributes += fmt.Sprintf(" %s=%q", attr.Key, attr.Val)
		}
		if n.FirstChild != nil {
			fmt.Fprintf(w, "%*s<%s%s>\n", depth*2, "", n.Data, attributes)
		} else {
			fmt.Fprintf(w, "%*s<%s%s />\n", depth*2, "", n.Data, attributes)
		}
		depth++
	}

	return depth
}

func endElement(w io.Writer, n *html.Node, depth int) int {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	return depth
}
