package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	length := len(s)
	intPartLength := strings.Index(s, ".")
	if intPartLength == -1 {
		intPartLength = length
	}
	cursor := intPartLength % 3
	if cursor > 0 {
		buf.WriteString(s[:cursor])
	}

	for cursor < intPartLength {
		if cursor > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[cursor : cursor+3])
		cursor += 3
	}

	if cursor == length {
		return buf.String()
	}

	buf.WriteString(".")
	cursor++

	for cursor < length-3 {
		buf.WriteString(s[cursor : cursor+3])
		buf.WriteString(",")
		cursor += 3
	}

	buf.WriteString(s[cursor:])

	return buf.String()
}
