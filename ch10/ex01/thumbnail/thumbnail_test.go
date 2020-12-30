package thumbnail_test

import (
	"os"
	"testing"

	"github.com/progfay/go-training/ch10/ex01/thumbnail"
)

func Test_ImageFile(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    string
		want  string
	}{
		{
			title: "jpg",
			in:    "data/image.jpg",
			want:  "data/image.thumb.jpg",
		},
		{
			title: "png",
			in:    "data/image.png",
			want:  "data/image.thumb.png",
		},
		{
			title: "gif",
			in:    "data/image.gif",
			want:  "data/image.thumb.gif",
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			got, err := thumbnail.ImageFile(testcase.in)
			if err != nil {
				t.Error(err)
				return
			}

			if testcase.want != got {
				t.Errorf("want %q, got %q", testcase.want, got)
				return
			}

			_, err = os.Stat(got)
			if err != nil {
				t.Error(err)
				return
			}

			err = os.Remove(got)
			if err != nil {
				t.Error(err)
				return
			}
		})
	}
}
