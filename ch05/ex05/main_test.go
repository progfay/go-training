package main_test

import (
	"os"
	"testing"

	"golang.org/x/net/html"

	count "github.com/progfay/go-training/ch05/ex05"
)

type Out struct {
	words  int
	images int
}

var testcases = []struct {
	title    string
	htmlfile string
	want     Out
}{
	{
		title:    "empty html",
		htmlfile: "html/empty.html",
		want: Out{
			words:  0,
			images: 0,
		},
	},
	{
		title:    "nested html",
		htmlfile: "html/nested.html",
		want: Out{
			words:  4,
			images: 4,
		},
	},
	{
		title:    "sibling html",
		htmlfile: "html/sibling.html",
		want: Out{
			words:  4,
			images: 4,
		},
	},
	{
		title:    "longtext html",
		htmlfile: "html/longtext.html",
		want: Out{
			words:  69,
			images: 0,
		},
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

			words, images := count.InternalCountWordsAndImages(doc)

			if words != testcase.want.words {
				t.Errorf("want.words %v, got.words %v", testcase.want.words, words)
			}

			if images != testcase.want.images {
				t.Errorf("want.images %v, got.images %v", testcase.want.images, images)
			}
		})
	}
}
