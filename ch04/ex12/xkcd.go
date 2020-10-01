package main

import (
	"encoding/json"
	"io"
	"os"
)

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func loadComics() ([]*Comic, error) {
	f, err := os.Open("./comics")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	comics := []*Comic{}
	decoder := json.NewDecoder(f)
	for {
		comic := &Comic{}
		if err := decoder.Decode(comic); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		comics = append(comics, comic)
	}
	return comics, nil
}
