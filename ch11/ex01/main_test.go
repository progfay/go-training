package main

import (
	"errors"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func Test_countRune(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    string
		want  struct {
			counts  map[rune]int
			utflen  [utf8.UTFMax + 1]int
			invalid int
			err     error
		}
	}{
		{
			title: "alphabet",
			in:    "abcabc",
			want: struct {
				counts  map[rune]int
				utflen  [utf8.UTFMax + 1]int
				invalid int
				err     error
			}{
				counts: map[rune]int{
					'a': 2,
					'b': 2,
					'c': 2,
				},
				utflen:  [utf8.UTFMax + 1]int{0, 6, 0, 0, 0},
				invalid: 0,
				err:     nil,
			},
		},
		{
			title: "japanese",
			in:    "あいうえお",
			want: struct {
				counts  map[rune]int
				utflen  [utf8.UTFMax + 1]int
				invalid int
				err     error
			}{
				counts: map[rune]int{
					'あ': 1,
					'い': 1,
					'う': 1,
					'え': 1,
					'お': 1,
				},
				utflen:  [utf8.UTFMax + 1]int{0, 0, 0, 5, 0},
				invalid: 0,
				err:     nil,
			},
		},
		{
			title: "emoji",
			in:    "👀👨‍👩‍👧‍👦",
			want: struct {
				counts  map[rune]int
				utflen  [utf8.UTFMax + 1]int
				invalid int
				err     error
			}{
				counts: map[rune]int{
					'\u200d': 3,
					'👀':      1,
					'👨':      1,
					'👩':      1,
					'👧':      1,
					'👦':      1,
				},
				utflen:  [utf8.UTFMax + 1]int{0, 0, 0, 3, 5},
				invalid: 0,
				err:     nil,
			},
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			reader := strings.NewReader(testcase.in)
			counts, utflen, invalid, err := countRune(reader)

			if !reflect.DeepEqual(counts, testcase.want.counts) {
				t.Errorf("want counts %#v, got %#v", testcase.want.counts, counts)
			}

			if !reflect.DeepEqual(utflen, testcase.want.utflen) {
				t.Errorf("want utflen %#v, got %#v", testcase.want.utflen, utflen)
			}

			if invalid != testcase.want.invalid {
				t.Errorf("want invalid %d, got %d", testcase.want.invalid, invalid)
			}

			if !errors.Is(err, testcase.want.err) {
				t.Errorf("want err %#v, got %#v", testcase.want.err, err)
			}
		})
	}
}
