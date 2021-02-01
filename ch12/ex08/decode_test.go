package sexpr

import (
	"bytes"
	"reflect"
	"testing"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func Test_Decoder(t *testing.T) {
	in := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	out, err := Marshal(in)

	got := Movie{}
	d := NewDecoder(bytes.NewReader(out))
	err = d.Decode(&got)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(in, got) {
		t.Errorf("want %#v, got %#v", in, got)
	}
}
