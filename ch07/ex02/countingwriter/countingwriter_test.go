package countingwriter_test

import (
	"bytes"
	"testing"

	"github.com/progfay/go-training/ch07/ex02/countingwriter"
)

func Test_CountingWriter(t *testing.T) {
	w, n := countingwriter.CountingWriter(&bytes.Buffer{})

	got := *n
	if got != 0 {
		t.Errorf("want %d, got %d", 0, got)
	}

	w.Write([]byte("0123456789"))
	got = *n
	if got != 10 {
		t.Errorf("want %d, got %d", 10, got)
	}

	w.Write([]byte("0123456789"))
	got = *n
	if got != 20 {
		t.Errorf("want %d, got %d", 20, got)
	}
}
