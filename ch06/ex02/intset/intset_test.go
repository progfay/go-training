package intset_test

import (
	"testing"

	"github.com/progfay/go-training/ch06/ex02/intset"
)

func Test_IntSet_AddAll(t *testing.T) {
	for _, testcase := range []struct {
		title    string
		in       []int
		occurErr bool
	}{
		{
			title:    "positive integer",
			in:       []int{1},
			occurErr: false,
		},
		{
			title:    "zero",
			in:       []int{0},
			occurErr: false,
		},
		{
			title:    "negative integer",
			in:       []int{-1},
			occurErr: true,
		},
		{
			title:    "duplication",
			in:       []int{2, 2},
			occurErr: false,
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			occurErr := false
			defer func() {
				if occurErr != testcase.occurErr {
					if occurErr {
						t.Error("expect no error is occurred, but error is occurred")
					} else {
						t.Error("expect error is occurred, but no error is occurred")
					}
				}
			}()

			defer func() {
				err := recover()
				if err != nil {
					occurErr = true
				}
			}()

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
