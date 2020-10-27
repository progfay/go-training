package eval_test

import (
	"math"
	"testing"

	"github.com/progfay/go-training/ch07/ex14/eval"
)

func Test_abs_Eval(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			expr string
			env  eval.Env
		}
		want float64
	}{
		{
			title: "literal",
			in: struct {
				expr string
				env  eval.Env
			}{
				expr: "|1|",
				env:  eval.Env{},
			},
			want: 1,
		},
		{
			title: "var",
			in: struct {
				expr string
				env  eval.Env
			}{
				expr: "|x|",
				env:  eval.Env{"x": -2},
			},
			want: 2,
		},
		{
			title: "unary",
			in: struct {
				expr string
				env  eval.Env
			}{
				expr: "|-3|",
				env:  eval.Env{},
			},
			want: 3,
		},
		{
			title: "binary",
			in: struct {
				expr string
				env  eval.Env
			}{
				expr: "|1 - 5|",
				env:  eval.Env{},
			},
			want: 4,
		},
		{
			title: "call",
			in: struct {
				expr string
				env  eval.Env
			}{
				expr: "|sin(rad)|",
				env:  eval.Env{"rad": math.Pi * 1.5},
			},
			want: 1,
		},
		{
			title: "multi-abs",
			in: struct {
				expr string
				env  eval.Env
			}{
				expr: "|-1| + |1|",
				env:  eval.Env{},
			},
			want: 2,
		},
		{
			title: "nest",
			in: struct {
				expr string
				env  eval.Env
			}{
				expr: "|(|-1|) - 2|",
				env:  eval.Env{},
			},
			want: 1,
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			expr, err := eval.Parse(testcase.in.expr)
			if err != nil {
				t.Errorf("invalid format of testcase: %s\n%v", testcase.in.expr, err)
				return
			}

			got := expr.Eval(testcase.in.env)

			if got != testcase.want {
				t.Errorf("want %#v, got %#v", testcase.want, got)
			}
		})
	}
}

func Test_abs_String(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    string
		want  string
	}{
		{
			title: "simple",
			in:    "|1|",
			want:  "|1.000000|",
		},
		{
			title: "complex",
			in:    "|1| + |x + 100| + |(|-1|) - 2| + |log(y)|",
			want:  "(((|1.000000| + |(x + 100.000000)|) + |(|(-1.000000)| - 2.000000)|) + |log(y)|)",
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			expr, err := eval.Parse(testcase.in)
			if err != nil {
				t.Errorf("invalid format of testcase: %s\n%v", testcase.in, err)
				return
			}

			got := expr.String()

			if got != testcase.want {
				t.Errorf("want %q, got %q", testcase.want, got)
			}
		})
	}
}
