package eval_test

import (
	"reflect"
	"testing"

	"github.com/progfay/go-training/ch07/ex13/eval"
)

func Test_tree_String(t *testing.T) {
	for _, testcase := range []string{
		"-1 - x",
		"-1 + -x",
		"5 / 9 * (F - 32)",
		"sqrt(A / pi)",
		"pow(x, 3) + pow(y, 3)",
	} {
		t.Run(testcase, func(t *testing.T) {
			want, err := eval.Parse(testcase)
			if err != nil {
				t.Errorf("invalid format of testcase: %s\n%v", testcase, err)
				return
			}

			str := want.String()
			got, err := eval.Parse(str)
			if err != nil {
				t.Errorf("invalid format of expr.String(): %s\n%v", str, err)
			}

			if !reflect.DeepEqual(want, got) {
				t.Errorf("want %#v, got %#v", want, got)
			}
		})
	}
}
