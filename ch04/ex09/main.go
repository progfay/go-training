package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	countMap := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		countMap[word]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("count\tword")
	for word, count := range countMap {
		if count < 50 {
			continue
		}
		fmt.Printf("%d\t%q\n", count, word)
	}
}
