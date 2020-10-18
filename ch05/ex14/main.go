package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func breadthFirst(f func(item, query string) []string, worklist []string, query string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item, query)...)
			}
		}
	}
}

func crawl(path, query string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Print(err)
		return []string{}
	}

	dirs := make([]string, 0)
	for _, file := range files {
		filename := file.Name()
		if strings.Contains(filename, query) {
			fmt.Println(filepath.Join(path, filename))
		}
		if file.IsDir() {
			dirs = append(dirs, filepath.Join(path, filename))
		}
	}

	return dirs
}

func main() {
	breadthFirst(crawl, []string{os.Args[1]}, os.Args[2])
}
