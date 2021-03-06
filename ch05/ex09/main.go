package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		s := Expand(fmt.Sprintf("$%[1]s : %[1]s", arg), numeronym)
		fmt.Println(s)
	}
}

func numeronym(s string) string {
	runes := []rune(s)
	return string(runes[0]) + strconv.Itoa(len(runes)-2) + string(runes[len(runes)-1])
}

func Expand(s string, f func(string) string) string {
	var (
		buf     bytes.Buffer
		isFirst bool = true
	)

	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if isFirst {
			isFirst = false
		} else {
			buf.WriteRune(' ')
		}

		text := scanner.Text()
		if strings.HasPrefix(text, "$") {
			buf.WriteString(f(text[1:]))
		} else {
			buf.WriteString(text)
		}
	}

	return buf.String()
}
