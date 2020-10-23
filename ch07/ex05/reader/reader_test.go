package reader_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/progfay/go-training/ch07/ex05/reader"
)

var testcases = []struct {
	title string
	in    struct {
		limit int64
		bytes []byte
	}
	want struct {
		n     int
		bytes []byte
		eof   bool
	}
}{
	{
		title: "limit 0",
		in: struct {
			limit int64
			bytes []byte
		}{
			limit: 0,
			bytes: []byte("1234567890"),
		},
		want: struct {
			n     int
			bytes []byte
			eof   bool
		}{
			n:     0,
			bytes: []byte(""),
			eof:   true,
		},
	},
	{
		title: "within limit",
		in: struct {
			limit int64
			bytes []byte
		}{
			limit: 100,
			bytes: []byte("1234567890"),
		},
		want: struct {
			n     int
			bytes []byte
			eof   bool
		}{
			n:     10,
			bytes: []byte("1234567890"),
			eof:   false,
		},
	},
	{
		title: "limit over",
		in: struct {
			limit int64
			bytes []byte
		}{
			limit: 5,
			bytes: []byte("1234567890"),
		},
		want: struct {
			n     int
			bytes []byte
			eof   bool
		}{
			n:     5,
			bytes: []byte("12345"),
			eof:   false,
		},
	},
}

func Test_reader(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			buf := bytes.NewBuffer(testcase.in.bytes)
			r := reader.LimitReader(buf, testcase.in.limit)

			b := make([]byte, len(testcase.in.bytes))
			n, err := r.Read(b)

			if (err != io.EOF) == testcase.want.eof {
				t.Error("unexpected EOF")
				return
			}

			if err != nil {
				if err == io.EOF {
					if !testcase.want.eof {
						t.Error("unexpected EOF")
					}
					return
				}
				t.Error(err)
				return
			}

			if testcase.want.eof {
				t.Error("EOF should occur on limitReader.Read")
				return
			}

			if n != testcase.want.n {
				t.Errorf("n: want %d, got %d", testcase.want.n, n)
			}

			if !bytes.Equal(b, testcase.want.bytes) {
				t.Errorf("bytes: want %v, got %v", testcase.want.bytes, b)
			}
		})
	}
}
