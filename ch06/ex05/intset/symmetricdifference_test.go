package intset_test

import (
	"testing"

	"github.com/progfay/go-training/ch06/ex05/intset"
)

func Test_IntSet_SymmetricDifference(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			left  []uint
			right []uint
		}
		want string
	}{
		{
			title: "empty sets",
			in: struct {
				left  []uint
				right []uint
			}{
				left:  []uint{},
				right: []uint{},
			},
			want: "{}",
		},
		{
			title: "no duplication",
			in: struct {
				left  []uint
				right []uint
			}{
				left:  []uint{1, 3, 5, 7},
				right: []uint{2, 4, 6, 8},
			},
			want: "{1 2 3 4 5 6 7 8}",
		},
		{
			title: "duplication",
			in: struct {
				left  []uint
				right []uint
			}{
				left:  []uint{1, 3, 5, 7},
				right: []uint{1, 3, 5, 7},
			},
			want: "{}",
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			left, right := &intset.IntSet{}, &intset.IntSet{}
			left.AddAll(testcase.in.left...)
			right.AddAll(testcase.in.right...)
			left.SymmetricDifference(right)
			got := left.String()

			if got != testcase.want {
				t.Errorf("want %s, got %s", testcase.want, got)
			}
		})
	}
}
