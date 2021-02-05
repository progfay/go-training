package equal

import (
	"math"
	"testing"
)

func TestEqual(t *testing.T) {
	for _, test := range []struct {
		x, y interface{}
		want bool
	}{
		{x: 1e-10, y: 2e-10, want: true},
		{x: 1e-9, y: 2e-9, want: false},
		{x: math.NaN(), y: math.NaN(), want: false},
		{x: math.Inf(0), y: math.Inf(0), want: false},
	} {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("Equal(%v, %v) = %t",
				test.x, test.y, !test.want)
		}
	}
}
