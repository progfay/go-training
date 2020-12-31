package tar

import (
	"archive/tar"
	"os"

	"github.com/progfay/go-training/ch10/ex02/archive"
)

func init() {
	archive.Register(archive.Loader{
		Name:  "tar",
		Match: Match,
		Load:  Load,
	})
}

func Match(infile string) bool {
	r, err := os.Open(infile)
	if err != nil {
		return false
	}

	tr := tar.NewReader(r)
	_, err = tr.Next()

	return err == nil
}

func Load(infile string) (interface{}, error) {
	r, err := os.Open(infile)
	if err != nil {
		return nil, err
	}

	return tar.NewReader(r), nil
}
