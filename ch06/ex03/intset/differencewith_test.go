package intset_test

import (
	"testing"

	"github.com/progfay/go-training/ch06/ex03/intset"
)

func Test_IntSet_UnionWith(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			left  []int
			right []int
		}
		want string
	}{
		{
			title: "empty sets",
			in: struct {
				left  []int
				right []int
			}{
				left:  []int{},
				right: []int{},
			},
			want: "{}",
		},
		{
			title: "no duplication",
			in: struct {
				left  []int
				right []int
			}{
				left:  []int{1, 3, 5, 7},
				right: []int{2, 4, 6, 8},
			},
			want: "{1 3 5 7}",
		},
		{
			title: "duplication",
			in: struct {
				left  []int
				right []int
			}{
				left:  []int{1, 3, 5, 7},
				right: []int{1, 3, 5, 7},
			},
			want: "{}",
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			left, right := &intset.IntSet{}, &intset.IntSet{}
			left.AddAll(testcase.in.left...)
			right.AddAll(testcase.in.right...)
			left.DifferenceWith(right)
			got := left.String()

			if got != testcase.want {
				t.Errorf("want %s, got %s", testcase.want, got)
			}
		})
	}
}
