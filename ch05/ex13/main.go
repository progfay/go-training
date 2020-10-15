package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/progfay/go-training/ch05/ex07/prettier"
	"github.com/progfay/go-training/ch05/ex13/links"
	"golang.org/x/net/html"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func getDoc(urlString string) (*html.Node, *url.URL, error) {
	resp, err := http.Get(urlString)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, nil, fmt.Errorf("getting %s: %s", urlString, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, nil, fmt.Errorf("parsing %s as HTML: %v", urlString, err)
	}

	return doc, resp.Request.URL, nil
}

func createFile(r io.Reader, path string) error {
	p := filepath.Join("dist", path)
	dir := filepath.Dir(p)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, r)
	if err != nil {
		return err
	}

	return nil
}

func crawl(link string) []string {
	doc, u, err := getDoc(link)
	if err != nil {
		log.Print(err)
		return []string{}
	}

	var buf bytes.Buffer
	prettier.WritePrettyHtml(&buf, doc)
	createFile(&buf, filepath.Join(u.Hostname(), u.Path))

	hrefs := links.Extract(doc)
	sameDomain := []string{}
	for _, href := range hrefs {
		l, err := u.Parse(href)
		if err != nil {
			continue
		}
		if u.Host == l.Host {
			sameDomain = append(sameDomain, l.String())
		}
	}
	return sameDomain
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
