package main

import (
	"fmt"
	"sort"

	"github.com/progfay/go-training/ch07/ex08/sorting"
	"github.com/progfay/go-training/ch07/ex08/track"
)

var tracks = []*track.Track{
	track.New("Go", "Delilah", "From the Roots Up", 2012, "3m38s"),
	track.New("Go", "Moby", "Moby", 1992, "3m37s"),
	track.New("Go Ahead", "Alicia Keys", "As I Am", 2007, "4m36s"),
	track.New("Ready 2 Go", "Martin Solveig", "Smash", 2011, "4m24s"),
}

func main() {
	fmt.Println("Order by title, year:")
	sort.Sort(sorting.Order(tracks, track.ByTitle, track.ByYear))
	track.PrintTracks(tracks)
}
