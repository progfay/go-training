package sexpr

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

type Token = interface{}
type Symbol struct{ name string }
type String struct{ value string }
type Int struct{ value int }
type Nil struct{}
type StartList struct{}
type EndList struct{}

type Decoder struct {
	r   io.Reader
	lex *lexer
}

func NewDecoder(r io.Reader) *Decoder {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(r)
	d := Decoder{
		r:   r,
		lex: lex,
	}
	return &d
}

func (d *Decoder) Token() (Token, error) {
	d.lex.next()
	switch d.lex.token {
	case scanner.EOF:
		return nil, io.EOF

	case scanner.Ident:
		text := d.lex.text()
		if text == "nil" {
			return Nil{}, nil
		}
		return Symbol{name: text}, nil

	case scanner.String:
		s, _ := strconv.Unquote(d.lex.text())
		return String{value: s}, nil

	case scanner.Int:
		i, _ := strconv.Atoi(d.lex.text())
		return Int{value: i}, nil

	case '(':
		return StartList{}, nil

	case ')':
		return EndList{}, nil
	}
	return nil, fmt.Errorf("unexpected token %q", d.lex.text())
}

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}
