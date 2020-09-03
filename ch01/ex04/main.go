package main

import (
	"bufio"
	"fmt"
	"os"
)

type Counter struct {
	Count int64
	Files []string
}

func (c *Counter) Add(file string) {
	c.Count++
	for _, f := range c.Files {
		if f == file {
			return
		}
	}
	c.Files = append(c.Files, file)
}

func main() {
	counters := make(map[string]*Counter)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines("stdin", os.Stdin, counters)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(arg, f, counters)
			f.Close()
		}
	}

	for line, counter := range counters {
		if counter.Count > 1 {
			fmt.Printf("%d\t%s\t%v\n", counter.Count, line, counter.Files)
		}
	}
}

func countLines(fileName string, f *os.File, counters map[string]*Counter) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counter, ok := counters[line]
		if !ok {
			counters[line] = &Counter{
				Count: 1,
				Files: []string{fileName},
			}
		} else {
			counter.Add(fileName)
		}
	}
}
