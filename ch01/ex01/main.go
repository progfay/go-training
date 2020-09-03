package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
