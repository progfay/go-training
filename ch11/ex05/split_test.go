package split_test

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			s   string
			sep string
		}
		want []string
	}{
		{
			title: "empty string",
			in: struct {
				s   string
				sep string
			}{
				s:   "",
				sep: ":",
			},
			want: []string{""},
		},
		{
			title: "s == sep",
			in: struct {
				s   string
				sep string
			}{
				s:   ":",
				sep: ":",
			},
			want: []string{"", ""},
		},
		{
			title: "split a:b:c by :",
			in: struct {
				s   string
				sep string
			}{
				s:   "a:b:c",
				sep: ":",
			},
			want: []string{"a", "b", "c"},
		},
		{
			title: "split ssepep by sep",
			in: struct {
				s   string
				sep string
			}{
				s:   "ssepep",
				sep: "sep",
			},
			want: []string{"s", "ep"},
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			got := strings.Split(testcase.in.s, testcase.in.sep)
			if !reflect.DeepEqual(testcase.want, got) {
				t.Errorf("want %#v, got %#v", testcase.want, got)
			}
		})
	}
}
