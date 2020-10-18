package max_test

import (
	"math"
	"testing"

	"github.com/progfay/go-training/ch05/ex15/max"
)

func Test_Max(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    []int64
		want  int64
	}{
		{
			title: "zero argument",
			in:    []int64{},
			want:  int64(math.MinInt64),
		},
		{
			title: "one argument",
			in:    []int64{1},
			want:  1,
		},
		{
			title: "many arguments",
			in:    []int64{1, 2, 3, 4, 5},
			want:  5,
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			out := max.Max(testcase.in...)
			if out != testcase.want {
				t.Errorf("want %d, got %d", testcase.want, out)
			}
		})
	}
}

func Test_MaxRequiresAtLeastOneArg(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			first int64
			rest  []int64
		}
		want int64
	}{
		{
			title: "one argument",
			in: struct {
				first int64
				rest  []int64
			}{
				first: 1,
				rest:  nil,
			},
			want: 1,
		},
		{
			title: "one argument",
			in: struct {
				first int64
				rest  []int64
			}{
				first: 1,
				rest:  []int64{2, 3, 4, 5},
			},
			want: 5,
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			out := max.MaxRequiresAtLeastOneArg(testcase.in.first, testcase.in.rest...)
			if out != testcase.want {
				t.Errorf("want %d, got %d", testcase.want, out)
			}
		})
	}
}
