package sexpr

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"text/scanner"
)

func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next()
	defer func() {

		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
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

func parseScientificFloat(s string) (float64, error) {
	pos := strings.Index(s, "e")

	var baseVal float64
	var expVal int64

	baseStr := s[0:pos]
	baseVal, err := strconv.ParseFloat(baseStr, 64)
	if err != nil {
		return 0, err
	}

	expStr := s[(pos + 1):]
	expVal, err = strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return baseVal * math.Pow10(int(expVal)), nil
}

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		text := lex.text()
		switch v.Kind() {
		case reflect.Bool:
			switch text {
			case "t":
				v.SetBool(true)

			case "nil":
				v.SetBool(false)

			default:
				panic(fmt.Sprintf("unexpected bool token %q", text))
			}
			lex.next()
			return

		default:
			if text == "nil" {
				v.Set(reflect.Zero(v.Type()))
				lex.next()
				return
			}
		}

	case scanner.String:
		s, err := strconv.Unquote(lex.text())
		if err != nil {
			panic(fmt.Sprintf("failed to unquote string: %q: %v", lex.text(), err))
		}
		v.SetString(s)
		lex.next()
		return

	case scanner.Int:
		i, err := strconv.Atoi(lex.text())
		if err != nil {
			panic(fmt.Sprintf("failed to parse int: %q: %v", lex.text(), err))
		}
		v.SetInt(int64(i))
		lex.next()
		return

	case scanner.Float:
		text := lex.text()
		f, err := parseScientificFloat(text)
		if err != nil {
			panic(fmt.Sprintf("failed to parse float: %q: %v", text, err))
		}
		v.SetFloat(f)
		lex.next()
		return

	case '(':
		lex.next()
		readList(lex, v)
		lex.next()
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice:
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct:
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}

	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}
