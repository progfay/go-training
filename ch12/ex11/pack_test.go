package params

import (
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type T struct {
	Bool   bool   `http:"boolean"`
	Int    int    `http:"integer"`
	String string `http:"string"`
	NoTag  string
}

func Test_Pack(t *testing.T) {
	in := T{
		Bool:   false,
		Int:    100,
		String: "str",
		NoTag:  "no tag",
	}

	out := Pack(in)

	want, err := url.ParseQuery("boolean=false&integer=100&string=str&notag=no+tag")
	if err != nil {
		t.Error(err)
		return
	}

	if !cmp.Equal(out.Query(), want) {
		t.Errorf(cmp.Diff(out.Query(), want))
	}
}
