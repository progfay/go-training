package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/progfay/go-training/ch02/ex02/lenconv"
)

func main() {
	if len(os.Args) <= 1 {
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		text := stdin.Text()
		l, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		m := lenconv.Meter(l)
		ft := lenconv.Feet(l)
		fmt.Printf("%s = %s, %s = %s\n",
			m, lenconv.MToFt(m), ft, lenconv.FtToM(ft))
	}

	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		m := lenconv.Meter(l)
		ft := lenconv.Feet(l)
		fmt.Printf("%s = %s, %s = %s\n",
			m, lenconv.MToFt(m), ft, lenconv.FtToM(ft))
	}
}
