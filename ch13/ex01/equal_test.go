package equal

import (
	"testing"
)

func TestEqual(t *testing.T) {
	for _, test := range []struct {
		x, y interface{}
		want bool
	}{
		{x: 1e-10, y: 2e-10, want: true},
		{x: 1e-9, y: 2e-9, want: false},
	} {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("Equal(%v, %v) = %t",
				test.x, test.y, !test.want)
		}
	}
}
