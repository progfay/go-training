package circular

import (
	"testing"
)

func Test_IsCircular(t *testing.T) {
	circularSlice := make([]interface{}, 1)
	circularSlice[0] = &circularSlice

	circularMap := make(map[bool]interface{})
	circularMap[true] = &circularMap

	circularStruct := struct{ v interface{} }{}
	circularStruct.v = &circularStruct

	for _, test := range []struct {
		in   interface{}
		want bool
	}{
		{in: 0, want: false},
		{in: "", want: false},
		{in: true, want: false},
		{in: [...]int{0, 1, 2}, want: false},
		{in: []int{0, 1, 2}, want: false},
		{in: map[int]bool{0: true, 1: false, 2: true}, want: false},
		{in: struct{ v int }{v: 2}, want: false},

		{in: circularSlice, want: true},
		{in: circularMap, want: true},
		{in: circularStruct, want: true},
	} {
		if IsCircular(test.in) != test.want {
			t.Errorf("IsCircular(%v) = %t", test.in, !test.want)
		}
	}
}
