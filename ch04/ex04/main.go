package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	rotate(a[:], 2)
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s)

	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		rotate(ints, 2)
		fmt.Printf("%v\n", ints)
	}
}

func rotate(s []int, n int) {
	a := append(s, s[:n]...)[n:]
	for i := 0; i < len(s); i++ {
		s[i] = a[i]
	}
}
