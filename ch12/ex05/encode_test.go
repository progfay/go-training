package sexpr

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_Marshal(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
		Flag            bool
	}
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
		Flag: true,
	}

	jsonBytes, err := Marshal(in)
	if err != nil {
		t.Error(err)
		return
	}

	var out Movie
	err = json.Unmarshal(jsonBytes, &out)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(in, out) {
		t.Errorf("want %#v, got %#v", in, out)
	}
}
