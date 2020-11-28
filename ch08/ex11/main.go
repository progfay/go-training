package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/progfay/go-training/ch08/ex11/fetch"
)

func main() {
	mode := flag.String("mode", "any", "fetch mode (supported: 'any', 'race')")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "at least one url argument is required")
		os.Exit(1)
	}

	var resp *http.Response
	var err error

	switch *mode {
	case "any":
		resp, err = fetch.Any(args)

	case "race":
		resp, err = fetch.Race(args)

	default:
		fmt.Fprintf(os.Stderr, "unsupported mode: %q", *mode)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	fmt.Println(resp.Request.URL)
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
