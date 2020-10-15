package main_test

import (
	"testing"

	topo "github.com/progfay/go-training/ch05/ex10"
)

var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},

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

var testcases = []struct {
	title string
	in    map[string]map[string]bool
}{
	{
		title: "empty",
		in:    make(map[string]map[string]bool),
	},
	{
		title: "prereqs",
		in:    prereqs,
	},
}

func Test_TopoSort(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			seen := make(map[string]bool)
			out := topo.TopoSort(testcase.in)
			for _, item := range out {
				for req := range testcase.in[item] {
					if !seen[req] {
						t.Errorf("invalid topological sorting: %q is required before %q", req, item)
						return
					}
				}
				seen[item] = true
			}
		})
	}
}
