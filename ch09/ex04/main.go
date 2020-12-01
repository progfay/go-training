package main

import (
	"flag"
	"strconv"
	"sync"

	"github.com/progfay/go-training/ch09/ex04/pipeline"
)

func main() {
	num := flag.Int("num", 10, "num of pipeline")
	flag.Parse()

	p := pipeline.New(*num, strconv.Itoa)

	var wg sync.WaitGroup
	wg.Add(*num)

	p.Send(func(string) {
		wg.Done()
	})

	wg.Wait()
}
