package sexpr

import (
	"testing"
)

func Test_Marshal(t *testing.T) {
	type Movie struct {
		Title    string
		Subtitle string
		Year     int
		Actor    map[string]string
		Oscars   []string
		Sequel   *string
	}
	in := Movie{
		Title:    "",
		Subtitle: "",
		Year:     0,
		Actor:    make(map[string]string),
		Oscars:   make([]string, 0),
	}

	jsonBytes, err := Marshal(in)
	if err != nil {
		t.Error(err)
		return
	}

	got := string(jsonBytes)
	want := `{"Actor":{},"Oscars":[]}`
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
