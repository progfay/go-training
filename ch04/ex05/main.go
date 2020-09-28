package main

import "fmt"

func main() {
	data := []string{"one", "one", "two", "three", "three", "three"}
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", removeDuplicateNeighbor(data))
}

func removeDuplicateNeighbor(strings []string) []string {
	a, b := 0, ""
	for i, s := range strings {
		if i == 0 || s != b {
			b = s
			strings[a] = s
			a++
		}
	}
	return strings[:a]
}
