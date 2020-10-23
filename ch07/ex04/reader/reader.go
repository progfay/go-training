package reader

import (
	"io"

	"golang.org/x/net/html"
)

type reader struct {
	p []byte
	i int
}

func NewReader(s string) *reader {
	r := reader{}
	r.p = []byte(s)
	r.i = 0
	return &r
}

func (r *reader) Read(p []byte) (n int, err error) {
	if r.i >= len(r.p) {
		return 0, io.EOF
	}
	n = copy(p, r.p[r.i:])
	r.i += n
	return
}

func main() {
	r := NewReader("")
	_, _ = html.Parse(r)
}
