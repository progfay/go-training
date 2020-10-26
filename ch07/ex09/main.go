package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"

	"github.com/progfay/go-training/ch07/ex08/sorting"
	"github.com/progfay/go-training/ch07/ex08/track"
)

var tableTemplate *template.Template

func init() {
	tableTemplate = template.Must(template.New("escape").Parse(`
	<table border="1" width="500" cellspacing="0" cellpadding="5">
		<tr>
			<th><a href="/?key=title">Title</a></th>
			<th><a href="/?key=artist">Artist</a></th>
			<th><a href="/?key=album">Album</a></th>
			<th><a href="/?key=year">Year</a></th>
			<th><a href="/?key=length">Length</a></th>
		</tr>

		{{range .Tracks}}
			<tr>
				<td>{{.Title}}</td>
				<td>{{.Artist}}</td>
				<td>{{.Album}}</td>
				<td>{{.Year}}</td>
				<td>{{.Length}}</td>
			</tr>
		{{end}}

	</table>`))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}

	tracks := []*track.Track{
		track.New("Go", "Delilah", "From the Roots Up", 2012, "3m38s"),
		track.New("Go", "Moby", "Moby", 1992, "3m37s"),
		track.New("Go Ahead", "Alicia Keys", "As I Am", 2007, "4m36s"),
		track.New("Ready 2 Go", "Martin Solveig", "Smash", 2011, "4m24s"),
	}

	if len(r.Form["key"]) > 0 && r.Form["key"][0] != "" {
		switch r.Form["key"][0] {
		case "title":
			sort.Sort(sorting.Order(tracks, track.ByTitle))

		case "article":
			sort.Sort(sorting.Order(tracks, track.ByArtist))

		case "album":
			sort.Sort(sorting.Order(tracks, track.ByAlbum))

		case "year":
			sort.Sort(sorting.Order(tracks, track.ByYear))

		case "length":
			sort.Sort(sorting.Order(tracks, track.ByLength))
		}
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tableTemplate.Execute(w, struct{ Tracks []*track.Track }{tracks}); err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listen on http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
