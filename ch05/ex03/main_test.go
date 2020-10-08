package main_test

import (
	"os"
	"sort"
	"testing"

	"golang.org/x/net/html"

	findlinks "github.com/progfay/go-training/ch05/ex03"
)

var testcases = []struct {
	title    string
	htmlfile string
	want     []string
}{
	{
		title:    "empty html",
		htmlfile: "html/empty.html",
		want:     []string{},
	},
	{
		title:    "nested html",
		htmlfile: "html/nested.html",
		want: []string{
			"buzz",
		},
	},
	{
		title:    "sibling html",
		htmlfile: "html/sibling.html",
		want: []string{
			"hoge",
			"fuga",
			"piyo",
		},
	},
	{
		title:    "script html",
		htmlfile: "html/script.html",
		want: []string{
			"paragraph",
		},
	},
	{
		title:    "style html",
		htmlfile: "html/style.html",
		want: []string{
			"paragraph",
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

			out := findlinks.Visit(nil, doc)

			if len(testcase.want) != len(out) {
				t.Errorf("wrong length: len(want) %v, len(got) %v", len(testcase.want), len(out))
				return
			}

			sort.Strings(testcase.want)
			sort.Strings(out)

			for i := 0; i < len(out); i++ {
				if testcase.want[i] != out[i] {
					t.Errorf("wrong length: want %v, got %v", testcase.want, out)
					return
				}
			}
		})
	}
}
