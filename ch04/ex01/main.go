package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	fmt.Println(countSha256Diff("x", "X"))
}

func countSha256Diff(s1, s2 string) (count int8) {
	b1 := sha256.Sum256([]byte(s1))
	b2 := sha256.Sum256([]byte(s2))

	for i := 0; i < 32; i++ {
		count += int8(pc[(b1[i] ^ b2[i])])
	}

	return count
}
