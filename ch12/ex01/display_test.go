package display

import (
	"bytes"
	"io"
	"os"
	"testing"
)

type S struct {
	a int
	b int
}

func Test_Display_StructKeyMap(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	stdout := os.Stdout
	os.Stdout = w

	m := make(map[S]int)
	m[S{a: 0, b: 0}] = 0
	m[S{a: 1, b: 1}] = 1

	Display("m", m)

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	if buf.String() != "Display m (map[display.S]int):\nm[{\"a\": 0, \"b\": 0}] = 0\nm[{\"a\": 1, \"b\": 1}] = 1\n" {
		t.Errorf("%q", buf.String())
	}
}

func Test_Display_ArrayKeyMap(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	stdout := os.Stdout
	os.Stdout = w

	m := make(map[[2]int]int)
	m[[2]int{0, 0}] = 0
	m[[2]int{1, 1}] = 1

	Display("m", m)

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	if buf.String() != "Display m (map[[2]int]int):\nm[[0, 0]] = 0\nm[[1, 1]] = 1\n" {
		t.Errorf("%q", buf.String())
	}
}
