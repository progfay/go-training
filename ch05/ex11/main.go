package main

import (
	"fmt"
	"os"
)

var ClosedCircuitError = fmt.Errorf("closed circuit was found")

var prereqs = map[string]map[string]bool{
	"algorithms":     {"data structures": true},
	"calculus":       {"linear algebra": true},
	"linear algebra": {"calculus": true},

	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

func main() {
	topo, err := TopoSort(prereqs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, course := range topo {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func copyMap(m map[string]bool) map[string]bool {
	copied := make(map[string]bool)
	for key, value := range m {
		copied[key] = value
	}
	return copied
}

func TopoSort(m map[string]map[string]bool) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items, visited map[string]bool) error

	visitAll = func(items, visited map[string]bool) error {
		for item := range items {
			if visited[item] {
				return ClosedCircuitError
			}

			if !seen[item] {
				visited[item] = true
				err := visitAll(m[item], copyMap(visited))
				if err != nil {
					return err
				}
				seen[item] = true
				order = append(order, item)
			}
		}
		return nil
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}

	err := visitAll(keys, make(map[string]bool))
	if err != nil {
		return nil, err
	}

	return order, nil
}
