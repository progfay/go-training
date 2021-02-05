package main_test

import (
	"strings"
	"testing"

	template "github.com/progfay/go-training/ch05/ex09"
)

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

var testcases = []struct {
	title string
	in    string
	fun   func(string) string
	want  string
}{
	{
		title: "empty",
		in:    "",
		fun:   func(s string) string { return "some-text" },
		want:  "",
	},
	{
		title: "no replace",
		in:    "aaa bbb ccc",
		fun:   func(s string) string { return "ddd" },
		want:  "aaa bbb ccc",
	},
	{
		title: "replace with empty text",
		in:    "$",
		fun: func(s string) string {
			if s == "" {
				return "empty-text"
			}
			return "some-text"
		},
		want: "empty-text",
	},
	{
		title: "single replace",
		in:    "abc $def ghe",
		fun:   reverse,
		want:  "abc fed ghe",
	},

	{
		title: "multi replace",
		in:    "$4321 $65 $0987",
		fun:   reverse,
		want:  "1234 56 7890",
	},
}

func Test_ElementByID(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			got := template.Expand(testcase.in, testcase.fun)
			want := strings.TrimSpace(testcase.want)
			if got != want {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}
