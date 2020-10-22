package main

import (
	"fmt"

	"github.com/progfay/go-training/ch07/ex01/counter"
)

func main() {
	var name = "Dolly"

	var wc counter.WordCounter
	wc.Write([]byte("one two three"))
	fmt.Println(wc)

	wc = 0
	fmt.Fprintf(&wc, "hello, %s", name)
	fmt.Println(wc)

	var lc counter.LineCounter
	lc.Write([]byte("one.\ntwo.\nthree."))
	fmt.Println(lc)

	lc = 0
	fmt.Fprintf(&lc, "hello, %s", name)
	fmt.Println(lc)
}
