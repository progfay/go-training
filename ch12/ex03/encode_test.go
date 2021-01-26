package sexpr

import "testing"

func Test_Marshal(t *testing.T) {
	for _, testcase := range []struct {
		name string
		in   interface{}
		want string
	}{
		{
			name: "truthy bool",
			in:   true,
			want: "t",
		},
		{
			name: "falsy bool",
			in:   false,
			want: "nil",
		},
		{
			name: "float",
			in:   12.345,
			want: "1.234500e+01",
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			b, err := Marshal(testcase.in)
			if err != nil {
				t.Error(err)
			}

			got := string(b)
			if got != testcase.want {
				t.Errorf("want %q, got %q", testcase.want, got)
			}
		})
	}
}
