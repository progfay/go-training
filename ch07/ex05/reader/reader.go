package reader

import (
	"io"
)

type limitReader struct {
	r io.Reader
	i int64
	l int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	lr := limitReader{}
	lr.r = r
	lr.i = 0
	lr.l = n
	return &lr
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	if lr.i >= lr.l {
		return 0, io.EOF
	}
	n, err = lr.r.Read(p)
	p = p[lr.i:lr.l]
	n = int(lr.l - lr.i)
	lr.i = lr.l
	return
}
