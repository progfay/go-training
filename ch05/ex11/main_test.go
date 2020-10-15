package main_test

import (
	"testing"

	topo "github.com/progfay/go-training/ch05/ex11"
)

var testcases = []struct {
	title    string
	in       map[string]map[string]bool
	isClosed bool
}{
	{
		title:    "empty",
		in:       map[string]map[string]bool{},
		isClosed: false,
	},
	{
		title: "not circuit",
		in: map[string]map[string]bool{
			"a": {"b": true},
			"b": {"c": true},
		},
		isClosed: false,
	},
	{
		title: "simple circuit",
		in: map[string]map[string]bool{
			"a": {"b": true},
			"b": {"a": true},
		},
		isClosed: true,
	},
	{
		title: "triangle circuit",
		in: map[string]map[string]bool{
			"a": {"b": true},
			"b": {"c": true},
			"c": {"a": true},
		},
		isClosed: true,
	},
	{
		title: "prereqs",
		in: map[string]map[string]bool{
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
		},
		isClosed: true,
	},
}

func Test_TopoSort(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			out, err := topo.TopoSort(testcase.in)

			if err != nil {
				if err == topo.ClosedCircuitError {
					switch testcase.isClosed {
					case true:
						return

					case false:
						t.Error("detected circuit when open circuit graph case")
						return
					}
				}

				t.Errorf("unexpected error: %v", err)
			}

			if testcase.isClosed {
				t.Error("not detected circuit when closed circuit case")
				return
			}

			seen := make(map[string]bool)
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
