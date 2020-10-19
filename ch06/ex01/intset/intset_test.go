package intset_test

import (
	"reflect"
	"testing"

	"github.com/progfay/go-training/ch06/ex01/intset"
)

func Test_IntSet_Has(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			entries []int
			arg     int
		}
		want bool
	}{
		{
			title: "empty",
			in: struct {
				entries []int
				arg     int
			}{
				entries: []int{},
				arg:     1,
			},
			want: false,
		},
		{
			title: "negative integer",
			in: struct {
				entries []int
				arg     int
			}{
				entries: []int{},
				arg:     -1,
			},
			want: false,
		},
		{
			title: "truly",
			in: struct {
				entries []int
				arg     int
			}{
				entries: []int{1},
				arg:     1,
			},
			want: true,
		},
		{
			title: "falsy",
			in: struct {
				entries []int
				arg     int
			}{
				entries: []int{1},
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
		title    string
		in       []int
		occurErr bool
	}{
		{
			title:    "positive integer",
			in:       []int{1},
			occurErr: false,
		},
		{
			title:    "zero",
			in:       []int{0},
			occurErr: false,
		},
		{
			title:    "negative integer",
			in:       []int{-1},
			occurErr: true,
		},
		{
			title:    "duplication",
			in:       []int{2, 2},
			occurErr: false,
		},
	} {
		t.Run(testcase.title, func(t *testing.T) {
			occurErr := false
			defer func() {
				if occurErr != testcase.occurErr {
					if occurErr {
						t.Error("expect no error is occurred, but error is occurred")
					} else {
						t.Error("expect error is occurred, but no error is occurred")
					}
				}
			}()

			defer func() {
				err := recover()
				if err != nil {
					occurErr = true
				}
			}()

			s := &intset.IntSet{}
			for _, v := range testcase.in {
				s.Add(v)
			}

			for _, v := range testcase.in {
				if !s.Has(v) {
					t.Errorf("added %[1]d to IntSet, but not have %[1]d", v)
				}
			}
		})
	}
}

func Test_IntSet_UnionWith(t *testing.T) {
	for _, testcase := range []struct {
		title string
		in    struct {
			left  []int
			right []int
		}
		want struct {
			have    []int
			notHave []int
		}
	}{
		{
			title: "union of empty sets",
			in: struct {
				left  []int
				right []int
			}{
				left:  []int{},
				right: []int{},
			},
			want: struct {
				have    []int
				notHave []int
			}{
				have:    []int{},
				notHave: []int{0, 1, 2, 3},
			},
		},
		{
			title: "no duplication",
			in: struct {
				left  []int
				right []int
			}{
				left:  []int{1, 3, 5, 7},
				right: []int{2, 4, 6, 8},
			},
			want: struct {
				have    []int
				notHave []int
			}{
				have:    []int{1, 2, 3, 4, 5, 6, 7, 8},
				notHave: []int{0},
			},
		},
		{
			title: "duplication",
			in: struct {
				left  []int
				right []int
			}{
				left:  []int{1, 3, 5, 7},
				right: []int{1, 3, 5, 7},
			},
			want: struct {
				have    []int
				notHave []int
			}{
				have:    []int{1, 3, 5, 7},
				notHave: []int{2, 4, 6, 8},
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
		in    []int
		want  string
	}{
		{
			title: "empty set",
			in:    []int{},
			want:  "{}",
		},
		{
			title: "set that have entries",
			in:    []int{1, 2, 3, 4, 5},
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
		in    []int
		want  int
	}{
		{
			title: "empty set",
			in:    []int{},
			want:  0,
		},
		{
			title: "no duplication",
			in:    []int{1, 2, 3, 4, 5},
			want:  5,
		},
		{
			title: "duplication",
			in:    []int{2, 2},
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
			entries []int
			arg     int
		}
	}{
		{
			title: "empty",
			in: struct {
				entries []int
				arg     int
			}{
				entries: []int{},
				arg:     1,
			},
		},
		{
			title: "remove exist integer",
			in: struct {
				entries []int
				arg     int
			}{
				entries: []int{1},
				arg:     1,
			},
		},
		{
			title: "remove non-exist integer",
			in: struct {
				entries []int
				arg     int
			}{
				entries: []int{1},
				arg:     0,
			},
		},
		{
			title: "negative integer",
			in: struct {
				entries []int
				arg     int
			}{
				entries: []int{},
				arg:     -1,
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
		in    []int
	}{
		{
			title: "empty",
			in:    []int{},
		},
		{
			title: "non-empty",
			in:    []int{1, 2, 3, 4},
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
		in    []int
	}{
		{
			title: "empty",
			in:    []int{},
		},
		{
			title: "non-empty",
			in:    []int{1, 2, 3, 4},
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
