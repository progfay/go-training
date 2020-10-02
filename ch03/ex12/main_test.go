package main_test

import (
	"testing"

	anagram "github.com/progfay/go-training/ch03/ex12"
)

var testcases = []struct {
	input [2]string
	want  bool
}{
	{input: [2]string{"", ""}, want: true},
	{input: [2]string{"abc", "acb"}, want: true},
	{input: [2]string{"acc", "abc"}, want: false},
	{input: [2]string{"🎉✨", "✨🎉"}, want: true},
}

func Test_IsAnagram(t *testing.T) {
	for _, testcase := range testcases {
		out := anagram.IsAnagram(testcase.input[0], testcase.input[1])
		if out != testcase.want {
			t.Errorf("IsAnagram(%q, %q), want %v, got %v", testcase.input[0], testcase.input[1], testcase.want, out)
		}
	}
}
