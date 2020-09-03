package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()

	fetch(os.Args[1])
	fmt.Fprintf(os.Stderr, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	nbytes, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Fprintf(os.Stderr, "%.2fs  %7d  %s", secs, nbytes, url)
}
