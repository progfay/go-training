package sorting_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/progfay/go-training/ch07/ex08/sorting"
	"github.com/progfay/go-training/ch07/ex08/track"
)

func Test_Order(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			tracks       []*track.Track
			compareFuncs []func(t1, t2 *track.Track) int
		}
		want []*track.Track
	}{
		{
			title: "no sorting",
			in: struct {
				tracks       []*track.Track
				compareFuncs []func(t1 *track.Track, t2 *track.Track) int
			}{
				tracks: []*track.Track{
					track.New("1", "1", "1", 1, "1s"),
					track.New("0", "0", "0", 0, "0s"),
					track.New("2", "2", "2", 2, "2s"),
				},
				compareFuncs: []func(t1 *track.Track, t2 *track.Track) int{},
			},
			want: []*track.Track{
				track.New("1", "1", "1", 1, "1s"),
				track.New("0", "0", "0", 0, "0s"),
				track.New("2", "2", "2", 2, "2s"),
			},
		},
		{
			title: "sort by artist",
			in: struct {
				tracks       []*track.Track
				compareFuncs []func(t1 *track.Track, t2 *track.Track) int
			}{
				tracks: []*track.Track{
					track.New("1", "1", "1", 1, "1s"),
					track.New("0", "0", "0", 0, "0s"),
					track.New("2", "2", "2", 2, "2s"),
				},
				compareFuncs: []func(t1 *track.Track, t2 *track.Track) int{track.ByArtist},
			},
			want: []*track.Track{
				track.New("0", "0", "0", 0, "0s"),
				track.New("1", "1", "1", 1, "1s"),
				track.New("2", "2", "2", 2, "2s"),
			},
		},
		{
			title: "sort by artist, year",
			in: struct {
				tracks       []*track.Track
				compareFuncs []func(t1 *track.Track, t2 *track.Track) int
			}{
				tracks: []*track.Track{
					track.New("1", "1", "1", 1, "1s"),
					track.New("0", "0", "0", 0, "0s"),
					track.New("1", "1", "1", 2, "1s"),
				},
				compareFuncs: []func(t1 *track.Track, t2 *track.Track) int{track.ByArtist, track.ByYear},
			},
			want: []*track.Track{
				track.New("0", "0", "0", 0, "0s"),
				track.New("1", "1", "1", 1, "1s"),
				track.New("1", "1", "1", 2, "1s"),
			},
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			sort.Sort(sorting.Order(testcase.in.tracks, testcase.in.compareFuncs...))

			if len(testcase.want) != len(testcase.in.tracks) {
				t.Errorf("len(want) %v, len(got) %v", len(testcase.want), len(testcase.in.tracks))
			}

			for i := 0; i < len(testcase.want); i++ {
				if !reflect.DeepEqual(*testcase.in.tracks[i], *testcase.want[i]) {
					t.Errorf("want %v, got %v", *testcase.want[i], *testcase.in.tracks[i])
				}
			}
		})
	}
}
