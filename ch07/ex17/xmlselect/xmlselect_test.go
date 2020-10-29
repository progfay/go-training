package xmlselect_test

import (
	"bytes"
	"io"
	"reflect"
	"sort"
	"testing"

	"github.com/progfay/go-training/ch07/ex17/xmlselect"
)

func Test_QuerySelectorAll(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			reader   io.Reader
			selector []string
		}
		want []string
	}{
		{
			title: "name",
			in: struct {
				reader   io.Reader
				selector []string
			}{
				reader: bytes.NewBufferString(`
					<a>
						<b>
							<c>match</c>
							<d>not match</d>
						</b>
					</a>
				`),
				selector: []string{"a", "b", "c"},
			},
			want: []string{"match"},
		},
		{
			title: "class",
			in: struct {
				reader   io.Reader
				selector []string
			}{
				reader: bytes.NewBufferString(`
					<a class="1">
						<b class="2">
							<c class="3">match</c>
							<c class="4">not match</c>
						</b>
					</a>
				`),
				selector: []string{".1", ".2", ".3"},
			},
			want: []string{"match"},
		},
		{
			title: "id",
			in: struct {
				reader   io.Reader
				selector []string
			}{
				reader: bytes.NewBufferString(`
					<a id="one">
						<b id="two">
							<c id="three">match</c>
							<c>not match</c>
						</b>
					</a>
				`),
				selector: []string{"#one", "#two", "#three"},
			},
			want: []string{"match"},
		},
		{
			title: "multiple matching",
			in: struct {
				reader   io.Reader
				selector []string
			}{
				reader: bytes.NewBufferString(`
					<a>
						<b>
							<c>match1</c>
							<c>match2</c>
						</b>
						<b>
							<c>match3</c>
						</b>
					</a>
					<a>
						<b>
							<c>match4</c>
						</b>
					</a>
				`),
				selector: []string{"a", "b", "c"},
			},
			want: []string{
				"match1",
				"match2",
				"match3",
				"match4",
			},
		},
		{
			title: "disturber",
			in: struct {
				reader   io.Reader
				selector []string
			}{
				reader: bytes.NewBufferString(`
					<a>
						<div>
							<b>
								<c>match</c>
							</b>
						</div>
					</a>
				`),
				selector: []string{"a", "b", "c"},
			},
			want: []string{"match"},
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			got := xmlselect.QuerySelectorAll(testcase.in.reader, testcase.in.selector)

			sort.Strings(got)
			sort.Strings(testcase.want)

			if !reflect.DeepEqual(got, testcase.want) {
				t.Errorf("want %v, got %v", testcase.want, got)
			}
		})
	}
}
