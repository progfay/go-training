package has_test

import (
	"testing"

	"github.com/progfay/go-training/ch10/ex04/has"
)

func Test_Duplicate(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			left  []string
			right []string
		}
		want bool
	}{
		{
			title: "empty slices",
			in: struct {
				left  []string
				right []string
			}{
				left:  []string{},
				right: []string{},
			},
			want: true,
		},
		{
			title: "empty slice only left",
			in: struct {
				left  []string
				right []string
			}{
				left:  []string{},
				right: []string{""},
			},
			want: false,
		},
		{
			title: "empty slice only right",
			in: struct {
				left  []string
				right []string
			}{
				left:  []string{""},
				right: []string{},
			},
			want: false,
		},
		{
			title: "has duplication",
			in: struct {
				left  []string
				right []string
			}{
				left:  []string{"one", "two", "three"},
				right: []string{"1", "two", "3"},
			},
			want: true,
		},
		{
			title: "perfect different",
			in: struct {
				left  []string
				right []string
			}{
				left:  []string{"one", "two", "three"},
				right: []string{"1", "2", "3"},
			},
			want: false,
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			got := has.Duplicate(testcase.in.left, testcase.in.right)

			if testcase.want != got {
				t.Errorf("want %t, got %t", testcase.want, got)
			}
		})
	}
}
