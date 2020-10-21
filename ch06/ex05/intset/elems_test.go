package intset_test

import (
	"reflect"
	"testing"

	"github.com/progfay/go-training/ch06/ex05/intset"
)

func Test_IntSet_Elems(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    []uint
		want  []uint
	}{
		{
			title: "empty set",
			in:    []uint{},
			want:  []uint{},
		},
		{
			title: "no duplication",
			in:    []uint{0, 1, 2},
			want:  []uint{0, 1, 2},
		},
		{
			title: "duplication",
			in:    []uint{2, 2},
			want:  []uint{2},
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			s := &intset.IntSet{}
			s.AddAll(testcase.in...)
			got := s.Elems()

			if !reflect.DeepEqual(got, testcase.want) {
				t.Errorf("want %v, got %v", testcase.want, got)
			}
		})
	}
}
