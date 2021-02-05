package sexpr

import (
	"bytes"
	"fmt"
	"io"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func Example_Unmarshal() {
	in := `((Title "Dr. Strangelove") (Subtitle "How I Learned to Stop Worrying and Love the Bomb") (Year 1964) (Actor (("Grp. Capt. Lionel Mandrake" "Peter Sellers") ("Pres. Merkin Muffley" "Peter Sellers") ("Gen. Buck Turgidson" "George C. Scott") ("Brig. Gen. Jack D. Ripper" "Sterling Hayden") ("Maj. T.J. \"King\" Kong" "Slim Pickens") ("Dr. Strangelove" "Peter Sellers"))) (Oscars ("Best Actor (Nomin.)" "Best Adapted Screenplay (Nomin.)" "Best Director (Nomin.)" "Best Picture (Nomin.)")) (Sequel nil))`
	d := NewDecoder(bytes.NewReader([]byte(in)))
	for {
		token, err := d.Token()
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		fmt.Printf("%#v\n", token)
	}

	// Output:
	// sexpr.StartList{}
	// sexpr.StartList{}
	// sexpr.Symbol{name:"Title"}
	// sexpr.String{value:"Dr. Strangelove"}
	// sexpr.EndList{}
	// sexpr.StartList{}
	// sexpr.Symbol{name:"Subtitle"}
	// sexpr.String{value:"How I Learned to Stop Worrying and Love the Bomb"}
	// sexpr.EndList{}
	// sexpr.StartList{}
	// sexpr.Symbol{name:"Year"}
	// sexpr.Int{value:1964}
	// sexpr.EndList{}
	// sexpr.StartList{}
	// sexpr.Symbol{name:"Actor"}
	// sexpr.StartList{}
	// sexpr.StartList{}
	// sexpr.String{value:"Grp. Capt. Lionel Mandrake"}
	// sexpr.String{value:"Peter Sellers"}
	// sexpr.EndList{}
	// sexpr.StartList{}
	// sexpr.String{value:"Pres. Merkin Muffley"}
	// sexpr.String{value:"Peter Sellers"}
	// sexpr.EndList{}
	// sexpr.StartList{}
	// sexpr.String{value:"Gen. Buck Turgidson"}
	// sexpr.String{value:"George C. Scott"}
	// sexpr.EndList{}
	// sexpr.StartList{}
	// sexpr.String{value:"Brig. Gen. Jack D. Ripper"}
	// sexpr.String{value:"Sterling Hayden"}
	// sexpr.EndList{}
	// sexpr.StartList{}
	// sexpr.String{value:"Maj. T.J. \"King\" Kong"}
	// sexpr.String{value:"Slim Pickens"}
	// sexpr.EndList{}
	// sexpr.StartList{}
	// sexpr.String{value:"Dr. Strangelove"}
	// sexpr.String{value:"Peter Sellers"}
	// sexpr.EndList{}
	// sexpr.EndList{}
	// sexpr.EndList{}
	// sexpr.StartList{}
	// sexpr.Symbol{name:"Oscars"}
	// sexpr.StartList{}
	// sexpr.String{value:"Best Actor (Nomin.)"}
	// sexpr.String{value:"Best Adapted Screenplay (Nomin.)"}
	// sexpr.String{value:"Best Director (Nomin.)"}
	// sexpr.String{value:"Best Picture (Nomin.)"}
	// sexpr.EndList{}
	// sexpr.EndList{}
	// sexpr.StartList{}
	// sexpr.Symbol{name:"Sequel"}
	// sexpr.Nil{}
	// sexpr.EndList{}
	// sexpr.EndList{}
}
