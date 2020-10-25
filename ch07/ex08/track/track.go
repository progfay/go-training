package track

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func New(title, artist, album string, year int, duration string) *Track {
	return &Track{
		Title:  title,
		Artist: artist,
		Album:  album,
		Year:   year,
		Length: length(duration),
	}
}

func ByTitle(t1, t2 *Track) int {
	switch {
	case t1.Title == t2.Title:
		return 0

	case t1.Title < t2.Title:
		return 1

	default:
		return -1
	}
}

func ByArtist(t1, t2 *Track) int {
	switch {
	case t1.Artist == t2.Artist:
		return 0

	case t1.Artist < t2.Artist:
		return 1

	default:
		return -1
	}
}

func ByAlbum(t1, t2 *Track) int {
	switch {
	case t1.Album == t2.Album:
		return 0

	case t1.Album < t2.Album:
		return 1

	default:
		return -1
	}
}

func ByYear(t1, t2 *Track) int {
	return t2.Year - t1.Year
}

func ByLength(t1, t2 *Track) int {
	return int(t2.Length - t1.Length)
}

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}
