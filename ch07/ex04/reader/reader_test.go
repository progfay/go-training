package reader_test

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"

	"github.com/progfay/go-training/ch07/ex04/reader"
)

var testcases = []struct {
	title string
	in    string
}{
	{
		title: "empty html",
		in:    "../html/empty.html",
	},
	{
		title: "nested html",
		in:    "../html/nested.html",
	},
	{
		title: "sibling html",
		in:    "../html/sibling.html",
	},
}

func Test_reader(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			f, err := os.Open(testcase.in)
			if err != nil {
				t.Errorf("fail on open %s: %v", testcase.in, err)
				return
			}

			b := make([]byte, 0)
			_, err = f.Read(b)
			if err != nil {
				t.Errorf("fail on read %s: %v", testcase.in, err)
				return
			}

			s := string(b)

			r := reader.NewReader(s)
			got, err := html.Parse(r)
			if err != nil {
				t.Errorf("fail on parse with reader.reader: %v", err)
				return
			}

			want, err := html.Parse(strings.NewReader(s))
			if err != nil {
				t.Errorf("fail on parse with strings.Reader: %v", err)
				return
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
				return
			}
		})
	}
}
