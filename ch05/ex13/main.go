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
	"syscall"

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

func curl(urlString string) (*bytes.Buffer, *url.URL, error) {
	resp, err := http.Get(urlString)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("getting %s: %s", urlString, resp.Status)
	}

	var buf bytes.Buffer
	io.Copy(&buf, resp.Body)

	return &buf, resp.Request.URL, nil
}

func parseToDoc(buf *bytes.Buffer) (*html.Node, error) {
	doc, err := html.Parse(buf)
	if err != nil {
		return nil, fmt.Errorf("fail html parsing: %v", err)
	}

	return doc, nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func createFile(r io.Reader, path string) error {
	p := filepath.Join("dist", path)
	if exists(p) {
		p = filepath.Join(p, "#")
	}
	dir := filepath.Dir(p)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		err, ok := err.(*os.PathError)
		if !ok || err.Err != syscall.ENOTDIR {
			return err
		}
		if err := os.Rename(err.Path, err.Path+"#"); err != nil {
			return err
		}
		if err := os.MkdirAll(err.Path, os.ModePerm); err != nil {
			return err
		}
		if err := os.Rename(err.Path+"#", filepath.Join(err.Path, "#")); err != nil {
			return err
		}
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
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
	buf, u, err := curl(link)
	if err != nil {
		log.Print(err)
		return []string{}
	}

	err = createFile(bytes.NewBuffer(buf.Bytes()), filepath.Join(u.Hostname(), u.Path))
	if err != nil {
		log.Print(err)
		return []string{}
	}

	doc, err := parseToDoc(bytes.NewBuffer(buf.Bytes()))
	if err != nil {
		log.Print(err)
		return []string{}
	}

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
