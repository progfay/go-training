package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Println(IsAnagram(os.Args[1], os.Args[2]))
}

func IsAnagram(str1, str2 string) bool {
	runes1, runes2 := []rune(str1), []rune(str2)
	if len(runes1) != len(runes2) {
		return false
	}

	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	return string(runes1) == string(runes2)
}
