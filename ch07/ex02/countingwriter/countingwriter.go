package countingwriter

import (
	"io"
)

type countingWriter struct {
	w io.Writer
	n int64
}

func new(w io.Writer) *countingWriter {
	c := countingWriter{}
	c.w = w
	c.n = 0
	return &c
}

func (c *countingWriter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.n += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := new(w)
	return c, &c.n
}
