package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	length := len(s)
	cursor := length % 3
	if cursor > 0 {
		buf.WriteString(s[:cursor])
	}

	for cursor < length {
		if cursor > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[cursor : cursor+3])
		cursor += 3
	}

	return buf.String()
}
