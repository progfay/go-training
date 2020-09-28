package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := []byte("123456789")
	reverse(a)
	fmt.Println(string(a))

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		bs := []byte(input.Text())
		reverse(bs)
		fmt.Println(string(bs))
	}
}

func reverse(bs []byte) {
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
}
