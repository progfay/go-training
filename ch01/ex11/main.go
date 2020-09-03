package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	client http.Client
)

func init() {
	client = http.Client{
		Timeout: 15 * time.Second,
	}
}

func main() {
	start := time.Now()
	ch := make(chan string)

	// file "top100" is extracted url list from https://moz.com/top500
	f, err := os.Open("top100")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	count := 0

	for input.Scan() {
		go fetch(input.Text(), ch)
		count++
	}
	for i := 0; i < count; i++ {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
