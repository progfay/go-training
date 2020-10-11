package main_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"

	"github.com/progfay/go-training/ch05/ex07/prettier"
	get "github.com/progfay/go-training/ch05/ex08"
)

var testcases = []struct {
	title    string
	htmlfile string
	id       string
	want     string
}{
	{
		title:    "empty html",
		htmlfile: "html/empty.html",
		id:       "empty",
		want:     "",
	},
	{
		title:    "not found",
		htmlfile: "html/not_found.html",
		id:       "found",
		want:     "",
	},
	{
		title:    "deep element",
		htmlfile: "html/deep_element.html",
		id:       "deep",
		want:     "<div id=\"deep\" />",
	},
	{
		title:    "shallow element",
		htmlfile: "html/shallow_element.html",
		id:       "shallow",
		want:     "<div id=\"shallow\" />",
	},
}

func Test_ElementByID(t *testing.T) {
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

			node := get.ElementByID(doc, testcase.id)

			var b bytes.Buffer
			prettier.WritePrettyHtml(&b, node)

			got := strings.TrimSpace(b.String())
			want := strings.TrimSpace(testcase.want)
			if got != want {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}
