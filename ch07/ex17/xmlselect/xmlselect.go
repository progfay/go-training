package xmlselect

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func QuerySelectorAll(r io.Reader, selector []string) []string {
	var (
		dec     = xml.NewDecoder(r)
		matched = make([]string, 0)
		stack   []xml.StartElement
	)

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if match(stack, selector) {
				matched = append(matched, string(tok))
			}
		}
	}

	return matched
}

func match(elements []xml.StartElement, selector []string) bool {
	var e xml.StartElement
	for len(selector) <= len(elements) {
		if len(selector) == 0 {
			return true
		}

		s := selector[0]
		e, elements = elements[0], elements[1:]
		match := false
		switch {
		case strings.HasPrefix(s, "."):
			s = s[1:]
			for _, attr := range e.Attr {
				if attr.Name.Local == "class" {
					match = attr.Value == s
					break
				}
			}

		case strings.HasPrefix(s, "#"):
			s = s[1:]
			for _, attr := range e.Attr {
				if attr.Name.Local == "id" {
					match = attr.Value == s
					break
				}
			}

		default:
			match = e.Name.Local == s
		}
		if match {
			selector = selector[1:]
		}
	}
	return false
}
