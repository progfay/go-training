package main_test

import (
	"os"
	"testing"

	"golang.org/x/net/html"

	findlinks "github.com/progfay/go-training/ch05/ex02"
)

var testcases = []struct {
	title    string
	htmlfile string
	want     map[string]int
}{
	{
		title:    "empty html",
		htmlfile: "html/empty.html",
		want: map[string]int{
			"body": 1,
			"head": 1,
			"html": 1,
		},
	},
	{
		title:    "nested html",
		htmlfile: "html/nested.html",
		want: map[string]int{
			"body": 1,
			"head": 1,
			"html": 1,
			"a":    3,
		},
	},
	{
		title:    "sibling html",
		htmlfile: "html/sibling.html",
		want: map[string]int{
			"body": 1,
			"head": 1,
			"html": 1,
			"span": 1,
			"a":    3,
		},
	},
}

func Test_visit(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			file, err := os.Open(testcase.htmlfile)
			if err != nil {
				t.Error(err)
				return
			}
			doc, err := html.Parse(file)
			if err != nil {
				t.Error(err)
				return
			}
			out := make(map[string]int)
			findlinks.Visit(out, doc)

			if len(testcase.want) != len(out) {
				t.Errorf("wrong length: len(want) %v, len(got) %v", len(testcase.want), len(out))
				return
			}

			for data, count := range testcase.want {
				if out[data] != count {
					t.Errorf("wrong length: want %v, got %v", testcase.want, out)
					return
				}
			}
		})
	}
}
