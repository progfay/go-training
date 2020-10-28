package eval_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/progfay/go-training/ch07/ex15/eval"
)

func Test_GetVars(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    string
		want  []string
	}{
		{
			title: "no vars",
			in:    "|1|",
			want:  []string{},
		},
		{
			title: "single var",
			in:    "x",
			want:  []string{"x"},
		},
		{
			title: "multiple vars",
			in:    "x + y + z",
			want:  []string{"x", "y", "z"},
		},
		{
			title: "duplicate",
			in:    "x + x + x",
			want:  []string{"x"},
		},
		{
			title: "nested",
			in:    "(a * b) + (c + (d + 4))",
			want:  []string{"a", "b", "c", "d"},
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			expr, err := eval.Parse(testcase.in)
			if err != nil {
				t.Errorf("invalid format of testcase: %s\n%v", testcase.in, err)
				return
			}

			got := eval.GetVars(expr)

			sort.Strings(got)
			sort.Strings(testcase.want)

			if !reflect.DeepEqual(got, testcase.want) {
				t.Errorf("want %v, got %v", testcase.want, got)
			}
		})
	}
}
