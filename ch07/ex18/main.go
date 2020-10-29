package main

import (
	"fmt"
	"os"

	"github.com/progfay/go-training/ch07/ex18/xmltree"
)

func main() {
	tree, err := xmltree.New(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", tree)
}
