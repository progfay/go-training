package main_test

import (
	"os"
	"sort"
	"testing"

	"golang.org/x/net/html"

	findlinks "github.com/progfay/go-training/ch05/ex04"
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
			"hoge",
			"fuga",
			"piyo",
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
		title:    "attributes html",
		htmlfile: "html/attributes.html",
		want: []string{
			"00",
			"01",
			"02",
			"03",
			"04",
			"05",
			"06",
			"07",
			"08",
			"09",
			"10",
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
				t.Errorf("wrong length: len(want) %v, len(got) %v", len(testcase.want), out)
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
