package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// Categories is the map of set of Unicode characters in category
// Ref. https://golang.org/pkg/unicode/#pkg-variables
var Categories = map[string]*unicode.RangeTable{
	"C":  unicode.C,
	"Cc": unicode.Cc,
	"Cf": unicode.Cf,
	"Co": unicode.Co,
	"Cs": unicode.Cs,
	"L":  unicode.L,
	"Ll": unicode.Ll,
	"Lm": unicode.Lm,
	"Lo": unicode.Lo,
	"Lt": unicode.Lt,
	"Lu": unicode.Lu,
	"M":  unicode.M,
	"Mc": unicode.Mc,
	"Me": unicode.Me,
	"Mn": unicode.Mn,
	"N":  unicode.N,
	"Nd": unicode.Nd,
	"Nl": unicode.Nl,
	"No": unicode.No,
	"P":  unicode.P,
	"Pc": unicode.Pc,
	"Pd": unicode.Pd,
	"Pe": unicode.Pe,
	"Pf": unicode.Pf,
	"Pi": unicode.Pi,
	"Po": unicode.Po,
	"Ps": unicode.Ps,
	"S":  unicode.S,
	"Sc": unicode.Sc,
	"Sk": unicode.Sk,
	"Sm": unicode.Sm,
	"So": unicode.So,
	"Z":  unicode.Z,
	"Zl": unicode.Zl,
	"Zp": unicode.Zp,
	"Zs": unicode.Zs,
}

func main() {
	categoryMap := make(map[string]int)
	var utflen [utf8.UTFMax + 1]int

	in := bufio.NewReader(os.Stdin)
read:
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break read
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			categoryMap["invalid"]++
			continue read
		}

		utflen[n]++

		for category, rangeTable := range Categories {
			if unicode.Is(rangeTable, r) {
				categoryMap[category]++
				continue read
			}
		}

		categoryMap["not_found"]++
	}

	fmt.Printf("class\tcount\n")
	for c, n := range categoryMap {
		fmt.Printf("%s\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
}
