package intset_test

import (
	"reflect"
	"testing"

	"github.com/progfay/go-training/ch06/ex05/intset"
)

func Test_IntSet_Has(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			entries []uint
			arg     uint
		}
		want bool
	}{
		{
			title: "empty",
			in: struct {
				entries []uint
				arg     uint
			}{
				entries: []uint{},
				arg:     1,
			},
			want: false,
		},
		{
			title: "truly",
			in: struct {
				entries []uint
				arg     uint
			}{
				entries: []uint{1},
				arg:     1,
			},
			want: true,
		},
		{
			title: "falsy",
			in: struct {
				entries []uint
				arg     uint
			}{
				entries: []uint{1},
				arg:     2,
			},
			want: false,
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			s := &intset.IntSet{}
			for _, entry := range testcase.in.entries {
				s.Add(entry)
			}
			got := s.Has(testcase.in.arg)
			if got != testcase.want {
				t.Errorf("want %v, got %v", testcase.want, got)
			}
		})
	}
}

func Test_IntSet_Add(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    []uint
	}{
		{
			title: "positive uinteger",
			in:    []uint{1},
		},
		{
			title: "zero",
			in:    []uint{0},
		},
		{
			title: "duplication",
			in:    []uint{2, 2},
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			s := &intset.IntSet{}
			for _, v := range testcase.in {
				s.Add(v)
			}

			for _, v := range testcase.in {
				if !s.Has(v) {
					t.Errorf("added %[1]d to uintSet, but not have %[1]d", v)
				}
			}
		})
	}
}

func Test_IntSet_UnionWith(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			left  []uint
			right []uint
		}
		want struct {
			have    []uint
			notHave []uint
		}
	}{
		{
			title: "union of empty sets",
			in: struct {
				left  []uint
				right []uint
			}{
				left:  []uint{},
				right: []uint{},
			},
			want: struct {
				have    []uint
				notHave []uint
			}{
				have:    []uint{},
				notHave: []uint{0, 1, 2, 3},
			},
		},
		{
			title: "no duplication",
			in: struct {
				left  []uint
				right []uint
			}{
				left:  []uint{1, 3, 5, 7},
				right: []uint{2, 4, 6, 8},
			},
			want: struct {
				have    []uint
				notHave []uint
			}{
				have:    []uint{1, 2, 3, 4, 5, 6, 7, 8},
				notHave: []uint{0},
			},
		},
		{
			title: "duplication",
			in: struct {
				left  []uint
				right []uint
			}{
				left:  []uint{1, 3, 5, 7},
				right: []uint{1, 3, 5, 7},
			},
			want: struct {
				have    []uint
				notHave []uint
			}{
				have:    []uint{1, 3, 5, 7},
				notHave: []uint{2, 4, 6, 8},
			},
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			main, sub := &intset.IntSet{}, &intset.IntSet{}
			for _, l := range testcase.in.left {
				main.Add(l)
			}
			for _, r := range testcase.in.right {
				main.Add(r)
			}
			main.UnionWith(sub)

			for _, v := range testcase.want.have {
				if !main.Has(v) {
					t.Errorf("union of set should have %[1]d, but not have %[1]d", v)
				}
			}

			for _, v := range testcase.want.notHave {
				if main.Has(v) {
					t.Errorf("union of set should not have %[1]d, but have %[1]d", v)
				}
			}
		})
	}
}

func Test_IntSet_String(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    []uint
		want  string
	}{
		{
			title: "empty set",
			in:    []uint{},
			want:  "{}",
		},
		{
			title: "set that have entries",
			in:    []uint{1, 2, 3, 4, 5},
			want:  "{1 2 3 4 5}",
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			s := &intset.IntSet{}
			for _, entry := range testcase.in {
				s.Add(entry)
			}

			got := s.String()
			if got != testcase.want {
				t.Errorf("want %q, got %q", testcase.want, got)
			}
		})
	}
}

func Test_IntSet_Len(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    []uint
		want  int
	}{
		{
			title: "empty set",
			in:    []uint{},
			want:  0,
		},
		{
			title: "no duplication",
			in:    []uint{1, 2, 3, 4, 5},
			want:  5,
		},
		{
			title: "duplication",
			in:    []uint{2, 2},
			want:  1,
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			s := &intset.IntSet{}
			for _, entry := range testcase.in {
				s.Add(entry)
			}

			got := s.Len()
			if got != testcase.want {
				t.Errorf("want %q, got %q", testcase.want, got)
			}
		})
	}
}

func Test_IntSet_Remove(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			entries []uint
			arg     uint
		}
	}{
		{
			title: "empty",
			in: struct {
				entries []uint
				arg     uint
			}{
				entries: []uint{},
				arg:     1,
			},
		},
		{
			title: "remove exist uinteger",
			in: struct {
				entries []uint
				arg     uint
			}{
				entries: []uint{1},
				arg:     1,
			},
		},
		{
			title: "remove non-exist uinteger",
			in: struct {
				entries []uint
				arg     uint
			}{
				entries: []uint{1},
				arg:     0,
			},
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			s := &intset.IntSet{}
			for _, entry := range testcase.in.entries {
				s.Add(entry)
			}

			s.Remove(testcase.in.arg)

			if s.Has(testcase.in.arg) {
				t.Errorf("remove %[1]d from set, but %[1]d still in set", testcase.in.arg)
			}
		})
	}
}

func Test_IntSet_Clear(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    []uint
	}{
		{
			title: "empty",
			in:    []uint{},
		},
		{
			title: "non-empty",
			in:    []uint{1, 2, 3, 4},
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			s := &intset.IntSet{}
			for _, entry := range testcase.in {
				s.Add(entry)
			}

			s.Clear()
			c := &intset.IntSet{}
			if !reflect.DeepEqual(s, c) {
				t.Errorf("cleared set is not empty: %#v %#v", s, c)
			}
		})
	}
}

func Test_IntSet_Copy(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    []uint
	}{
		{
			title: "empty",
			in:    []uint{},
		},
		{
			title: "non-empty",
			in:    []uint{1, 2, 3, 4},
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			s := &intset.IntSet{}
			for _, entry := range testcase.in {
				s.Add(entry)
			}

			copied := s.Copy()
			if !reflect.DeepEqual(s, copied) {
				t.Error("not equal deeply origin and copied")
			}

			if s == copied {
				t.Error("copied set has same address as origin")
			}
		})
	}
}
