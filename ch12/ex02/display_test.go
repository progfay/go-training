package display

import (
	"bytes"
	"io"
	"os"
	"testing"
)

type A struct{ t string }
type B struct {
	t string
	a A
}
type C struct {
	t string
	b B
}
type D struct {
	t string
	c C
}

func Test_Display_Recursive(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	stdout := os.Stdout
	os.Stdout = w

	d := D{
		t: "d",
		c: C{
			t: "c",
			b: B{
				t: "b",
				a: A{
					t: "a",
				},
			},
		},
	}
	Display("d", d)

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	if buf.String() != "Display d (display.D):\nd.t = \"d\"\nd.c.t = \"c\"\n" {
		t.Errorf("%q", buf.String())
	}
}
