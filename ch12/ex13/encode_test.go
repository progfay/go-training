package sexpr

import (
	"testing"
)

func Test_Marshal(t *testing.T) {
	type T struct {
		A string `expr:"hoge"`
		B string
	}

	in := T{
		A: "hoge",
		B: "fuga",
	}

	out, err := Marshal(in)
	if err != nil {
		t.Error(err)
		return
	}

	got := string(out)
	want := `((hoge "hoge") (B "fuga"))`
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
