package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Println(isAnagram(os.Args[1], os.Args[2]))
}

func isAnagram(str1, str2 string) bool {
	runes1, runes2 := []rune(str1), []rune(str2)
	if len(runes1) != len(runes2) {
		return false
	}

	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	for i, r := range runes1 {
		if r != runes2[i] {
			return false
		}
	}

	return true
}
