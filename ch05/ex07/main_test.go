package main_test

import (
	"bytes"
	"os"
	"testing"

	"golang.org/x/net/html"

	"github.com/progfay/go-training/ch05/ex07/prettier"
)

var testcases = []struct {
	title    string
	htmlfile string
}{
	{
		title:    "empty html",
		htmlfile: "html/empty.html",
	},
	{
		title:    "comment html",
		htmlfile: "html/comment.html",
	},
	{
		title:    "text html",
		htmlfile: "html/text.html",
	},
	{
		title:    "self closing html",
		htmlfile: "html/self_closing.html",
	},
}

func Test_countWordsAndImages(t *testing.T) {
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

			var b bytes.Buffer
			prettier.WritePrettyHtml(&b, doc)

			_, err = html.Parse(&b)
			if err != nil {
				t.Errorf("error on parse prettied html: %s", err)
			}
		})
	}
}
