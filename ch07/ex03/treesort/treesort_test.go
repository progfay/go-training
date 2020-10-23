package treesort_test

import (
	"testing"

	"github.com/progfay/go-training/ch07/ex03/treesort"
)

func Test_tree_String(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    []int
		want  string
	}{
		{
			title: "empty",
			in:    []int{},
			want:  "",
		},
		{
			title: "biased to right",
			in:    []int{0, 1, 2, 3, 4},
			want:  "(0\\(1\\(2\\(3\\(4)))))",
		},
		{
			title: "biased to right",
			in:    []int{4, 3, 2, 1, 0},
			want:  "(((((0)/1)/2)/3)/4)",
		},
		{
			title: "balanced",
			in:    []int{3, 1, 0, 2, 5, 4, 6},
			want:  "(((0)/1\\(2))/3\\((4)/5\\(6)))",
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			tree := treesort.Sort(testcase.in)
			got := tree.String()

			if got != testcase.want {
				t.Errorf("want %q, got %q", testcase.want, got)
			}
		})
	}
}
