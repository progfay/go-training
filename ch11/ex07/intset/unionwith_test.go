package intset_test

import (
	"reflect"
	"testing"

	"github.com/progfay/go-training/ch11/ex02/intmap"
	"github.com/progfay/go-training/ch11/ex02/intset"
)

func Test_UnionWith(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			left  []int
			right []int
		}
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
		},
		{
			title: "has duplication",
			in: struct {
				left  []int
				right []int
			}{
				left:  []int{0, 1, 2, 3},
				right: []int{2, 3, 4, 5},
			},
		},
		{
			title: "no duplication",
			in: struct {
				left  []int
				right []int
			}{
				left:  []int{0, 1, 2, 3},
				right: []int{4, 5, 6, 7},
			},
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			sl := intset.New()
			sl.AddAll(testcase.in.left...)
			sr := intset.New()
			sr.AddAll(testcase.in.right...)
			sl.UnionWith(sr)
			s := sl.Elems()

			ml := intmap.New()
			ml.AddAll(testcase.in.left...)
			mr := intmap.New()
			mr.AddAll(testcase.in.right...)
			ml.UnionWith(mr)
			m := ml.Elems()

			if !reflect.DeepEqual(s, m) {
				t.Errorf("intset %v, intmap %v", s, m)
			}
		})
	}
}
