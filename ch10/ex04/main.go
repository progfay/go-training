package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/progfay/go-training/ch10/ex04/list"
	"golang.org/x/sync/errgroup"
)

func main() {
	var (
		eg      errgroup.Group
		depPkgs []list.Package
		allPkgs []list.Package
	)

	flag.Parse()
	arg := flag.Arg(0)
	if arg == "" {
		fmt.Println("first argument packages must be required")
		os.Exit(1)
	}

	eg.Go(func() error {
		pkgs, err := list.Get(flag.Arg(0))
		if err != nil {
			return err
		}
		depPkgs = pkgs
		return nil
	})

	eg.Go(func() error {
		pkgs, err := list.Get("...")
		if err != nil {
			return err
		}
		allPkgs = pkgs
		return nil
	})

	err := eg.Wait()
	if err != nil {
		log.Panic(err)
	}

	deps := make([]string, len(depPkgs))
	for i, d := range depPkgs {
		deps[i] = d.ImportPath
	}
	sort.Strings(deps)

	for _, pkg := range allPkgs {
		sort.Strings(pkg.Deps)
		if hasDuplicate(deps, pkg.Deps) {
			fmt.Println(pkg.ImportPath)
		}
	}
}

// hasDuplicate detect duplication of string slice
// Arguments slices must be sorted
func hasDuplicate(left, right []string) bool {
	var (
		ll, rl = len(left), len(right)
		li, ri = 0, 0
	)

	for {
		if ll <= li || rl <= ri {
			return false
		}

		l, r := left[li], right[ri]
		switch {
		case l == r:
			return true

		case l < r:
			li++

		case l > r:
			ri++
		}
	}
}
