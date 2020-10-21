package intset_test

import (
	"testing"

	"github.com/progfay/go-training/ch06/ex05/intset"
)

func Test_IntSet_AddAll(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    []uint
	}{
		{
			title: "positive integer",
			in:    []uint{1},
		},
		{
			title: "zero",
			in:    []uint{0},
		},
		{
			title: "duplication",
			in:    []uint{2, 2},
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			s := &intset.IntSet{}
			s.AddAll(testcase.in...)

			for _, v := range testcase.in {
				if !s.Has(v) {
					t.Errorf("added %[1]d to IntSet, but not have %[1]d", v)
				}
			}
		})
	}
}
