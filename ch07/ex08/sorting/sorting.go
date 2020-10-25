package sorting

import (
	"github.com/progfay/go-training/ch07/ex08/track"
)

type customSort struct {
	tracks       []*track.Track
	compareFuncs []func(i1, i2 *track.Track) int
}

func (x customSort) Len() int {
	return len(x.tracks)
}

func (x customSort) Less(i, j int) bool {
	for _, compare := range x.compareFuncs {
		c := compare(x.tracks[i], x.tracks[j])
		if c == 0 {
			continue
		}
		return c > 0
	}
	return false
}

func (x customSort) Swap(i, j int) {
	x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i]
}

func Order(t []*track.Track, compareFuncs ...func(i1, i2 *track.Track) int) customSort {
	c := customSort{}
	c.tracks = t
	c.compareFuncs = compareFuncs
	return c
}
