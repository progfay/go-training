package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go func(root string) {
			defer wg.Done()
			du(root)
		}(root)
	}

	wg.Wait()
}

func du(root string) {
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	n.Add(1)
	go walkDir(root, &n, fileSizes)
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64

	for {
		size, ok := <-fileSizes
		if !ok {
			break
		}
		nfiles++
		nbytes += size
	}

	printDiskUsage(root, nfiles, nbytes)
}

func printDiskUsage(path string, nfiles, nbytes int64) {
	fmt.Printf("%q: %d files  %.1f GB\n", path, nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
