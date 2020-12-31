package zip

import (
	"archive/zip"
	"bufio"
	"os"

	"github.com/progfay/go-training/ch10/ex02/archive"
)

const (
	magic = "PK"
)

func init() {
	archive.Register(archive.Loader{
		Name:  "zip",
		Match: Match,
		Load:  Load,
	})
}

func Match(infile string) bool {
	r, err := os.Open(infile)
	if err != nil {
		return false
	}

	rr := bufio.NewReader(r)
	b, err := rr.Peek(len(magic))
	if err != nil {
		return false
	}

	return string(b) == magic
}

func Load(infile string) (interface{}, error) {
	return zip.OpenReader(infile)
}
