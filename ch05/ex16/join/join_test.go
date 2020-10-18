package join_test

import (
	"testing"

	"github.com/progfay/go-training/ch05/ex16/join"
)

func Test_Join(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			sep  string
			elms []string
		}
		want string
	}{
		{
			title: "zero elm",
			in: struct {
				sep  string
				elms []string
			}{
				sep:  " ",
				elms: []string{},
			},
			want: "",
		},
		{
			title: "one elm",
			in: struct {
				sep  string
				elms []string
			}{
				sep:  " ",
				elms: []string{"one"},
			},
			want: "one",
		},
		{
			title: "many elms",
			in: struct {
				sep  string
				elms []string
			}{
				sep:  " ",
				elms: []string{"one", "two", "three", "four", "five"},
			},
			want: "one two three four five",
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			out := join.Join(testcase.in.sep, testcase.in.elms...)
			if out != testcase.want {
				t.Errorf("want %s, got %s", testcase.want, out)
			}
		})
	}
}
