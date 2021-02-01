package sexpr

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type T struct {
	Truthy  bool
	Falsy   bool
	Float32 float32
	Float64 float64
}

func Test_Unmarshal(t *testing.T) {
	in := T{
		Truthy:  true,
		Falsy:   false,
		Float32: 123.456,
		Float64: 123.456,
	}

	out, err := Marshal(in)
	if err != nil {
		t.Error(err)
		return
	}

	got := T{}
	err = Unmarshal(out, &got)
	if err != nil {
		t.Error(err)
		return
	}

	opt := cmp.Comparer(func(x, y float64) bool {
		delta := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		return delta/mean < 0.00001
	})

	if !cmp.Equal(in, got, opt) {
		t.Errorf(cmp.Diff(in, got, opt))
	}
}
