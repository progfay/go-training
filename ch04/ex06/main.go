package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "You\tcan\r\rfly!"
	fmt.Printf("%q\n", str)
	fmt.Printf("%q\n", string(convertUnicodeSpaceToASCIISpace([]byte(str))))
}

func convertUnicodeSpaceToASCIISpace(bs []byte) []byte {
	a, f := 0, false
	for _, b := range bs {
		isSpace := unicode.IsSpace(rune(b))

		if !isSpace {
			bs[a] = b
			a++
		} else if !f {
			bs[a] = byte(' ')
			a++
		}

		f = isSpace
	}
	return bs[:a]
}
