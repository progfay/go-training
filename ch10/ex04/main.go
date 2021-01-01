package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/progfay/go-training/ch10/ex04/has"
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
		if has.Duplicate(deps, pkg.Deps) {
			fmt.Println(pkg.ImportPath)
		}
	}
}
