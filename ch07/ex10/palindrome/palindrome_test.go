package palindrome_test

import (
	"sort"
	"testing"

	"github.com/progfay/go-training/ch07/ex08/sorting"
	"github.com/progfay/go-training/ch07/ex08/track"
	"github.com/progfay/go-training/ch07/ex10/palindrome"
)

func Test_IsPalindrome(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    sort.Interface
		want  bool
	}{
		{
			title: "not palindrome int slice",
			in:    sort.IntSlice([]int{0, 1, 2, 3}),
			want:  false,
		},
		{
			title: "palindrome int slice (length: odd)",
			in:    sort.IntSlice([]int{0, 1, 0}),
			want:  true,
		},
		{
			title: "palindrome int slice (length: even)",
			in:    sort.IntSlice([]int{0, 1, 1, 0}),
			want:  true,
		},
		{
			title: "tracks (deep equal)",
			in: sorting.Order(
				[]*track.Track{
					track.New("0", "0", "0", 0, "0s"),
					track.New("1", "1", "1", 1, "1s"),
					track.New("0", "0", "0", 0, "0s"),
				},
				track.ByTitle,
			),
			want: true,
		},
		{
			title: "tracks (not deep equal)",
			in: sorting.Order(
				[]*track.Track{
					track.New("0", "0", "0", 0, "0s"),
					track.New("1", "1", "1", 1, "1s"),
					track.New("0", "2", "2", 2, "2s"),
				},
				track.ByTitle,
			),
			want: true,
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			got := palindrome.IsPalindrome(testcase.in)

			if testcase.want != got {
				t.Errorf("want %v, got %v", testcase.want, got)
			}
		})
	}
}
