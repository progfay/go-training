package sexpr

import (
	"reflect"
	"testing"
)

func Test_Unmarshal(t *testing.T) {
	type T struct {
		A string `expr:"hoge"`
		B string
	}

	in := `((hoge "hoge") (B "fuga"))`

	got := T{}
	err := Unmarshal([]byte(in), &got)
	if err != nil {
		t.Error(err)
		return
	}

	want := T{
		A: "hoge",
		B: "fuga",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %#v, got %#v", want, got)
	}
}
