package main

import (
	"os"

	"github.com/progfay/go-training/ch07/ex17/xmlselect"
)

func main() {
	xmlselect.QuerySelectorAll(os.Stdin, os.Args[1:])
}
