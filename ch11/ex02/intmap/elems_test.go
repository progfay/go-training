package intmap_test

import (
	"reflect"
	"testing"

	"github.com/progfay/go-training/ch11/ex02/intmap"
)

func Test_IntMap_Elems(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    []int
		want  []int
	}{
		{
			title: "empty set",
			in:    []int{},
			want:  []int{},
		},
		{
			title: "no duplication",
			in:    []int{0, 1, 2},
			want:  []int{0, 1, 2},
		},
		{
			title: "duplication",
			in:    []int{2, 2},
			want:  []int{2},
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			s := intmap.New()
			s.AddAll(testcase.in...)
			got := s.Elems()

			if !reflect.DeepEqual(got, testcase.want) {
				t.Errorf("want %v, got %v", testcase.want, got)
			}
		})
	}
}
