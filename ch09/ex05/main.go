package main

import (
	"fmt"
	"time"

	"github.com/progfay/go-training/ch09/ex05/twin"
)

func main() {
	t := twin.New()
	t.Run()
	time.Sleep(time.Second)
	fmt.Println(t.Count())
}
