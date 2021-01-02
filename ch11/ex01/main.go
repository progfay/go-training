package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func countRune(reader io.Reader) (counts map[rune]int, utflen [utf8.UTFMax + 1]int, invalid int, err error) {
	var (
		r rune
		n int
	)
	br := bufio.NewReader(reader)
	counts = make(map[rune]int)

	for {
		r, n, err = br.ReadRune()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	return
}

func main() {
	counts, utflen, invalid, err := countRune(os.Stdin)

	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
