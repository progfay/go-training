package main

import (
	"fmt"
	"os"
	"strings"

	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("require to set the search query as argument")
		os.Exit(1)
	}
	query := os.Args[1]

	comics, err := loadComics()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	width, err := terminal.Width()
	if err != nil {
		width = 30
	}

	for _, comic := range comics {
		if strings.Contains(comic.Title, query) {
			link := fmt.Sprintf("https://xkcd.com/%d", comic.Num)
			fmt.Printf("Title     : %s\nLink      : %s\nTranscript:\n%s\n", comic.Title, link, comic.Transcript)
			fmt.Println()
			fmt.Println(strings.Repeat("=", int(width)))
			fmt.Println()
		}
	}
}
