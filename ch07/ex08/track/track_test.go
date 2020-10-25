package track_test

import (
	"testing"

	"github.com/progfay/go-training/ch07/ex08/track"
)

var testcases = []struct {
	title string
	in    struct {
		left  *track.Track
		right *track.Track
	}
	want int
}{
	{
		title: "same",
		in: struct {
			left  *track.Track
			right *track.Track
		}{
			left:  track.New("Go", "Moby", "Moby", 2000, "0s"),
			right: track.New("Go", "Moby", "Moby", 2000, "0s"),
		},
		want: 0,
	},
	{
		title: "left is smaller than right",
		in: struct {
			left  *track.Track
			right *track.Track
		}{
			left:  track.New("Go", "Moby", "Moby", 2000, "0s"),
			right: track.New("Yo", "Noby", "Noby", 2001, "1s"),
		},
		want: 1,
	},
	{
		title: "left is bigger than right",
		in: struct {
			left  *track.Track
			right *track.Track
		}{
			left:  track.New("Yo", "Noby", "Noby", 2001, "1s"),
			right: track.New("Go", "Moby", "Moby", 2000, "0s"),
		},
		want: -1,
	},
}

func Test_ByTitle(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			got := track.ByTitle(testcase.in.left, testcase.in.right)

			switch testcase.want {
			case -1:
				if got >= 0 {
					t.Errorf("want negative value, got %d", got)
				}

			case 0:
				if got != 0 {
					t.Errorf("want 0, got %d", got)
				}

			case 1:
				if got <= 0 {
					t.Errorf("want positive value, got %d", got)
				}

			default:
				t.Errorf("Invalid direction value: %d", testcase.want)
			}
		})
	}
}

func Test_ByArtist(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			got := track.ByArtist(testcase.in.left, testcase.in.right)

			switch testcase.want {
			case -1:
				if got >= 0 {
					t.Errorf("want negative value, got %d", got)
				}

			case 0:
				if got != 0 {
					t.Errorf("want 0, got %d", got)
				}

			case 1:
				if got <= 0 {
					t.Errorf("want positive value, got %d", got)
				}

			default:
				t.Errorf("Invalid direction value: %d", testcase.want)
			}
		})
	}
}

func Test_ByAlbum(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			got := track.ByAlbum(testcase.in.left, testcase.in.right)

			switch testcase.want {
			case -1:
				if got >= 0 {
					t.Errorf("want negative value, got %d", got)
				}

			case 0:
				if got != 0 {
					t.Errorf("want 0, got %d", got)
				}

			case 1:
				if got <= 0 {
					t.Errorf("want positive value, got %d", got)
				}

			default:
				t.Errorf("Invalid direction value: %d", testcase.want)
			}
		})
	}
}

func Test_ByYear(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			got := track.ByYear(testcase.in.left, testcase.in.right)

			switch testcase.want {
			case -1:
				if got >= 0 {
					t.Errorf("want negative value, got %d", got)
				}

			case 0:
				if got != 0 {
					t.Errorf("want 0, got %d", got)
				}

			case 1:
				if got <= 0 {
					t.Errorf("want positive value, got %d", got)
				}

			default:
				t.Errorf("Invalid direction value: %d", testcase.want)
			}
		})
	}
}

func Test_ByLength(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			got := track.ByLength(testcase.in.left, testcase.in.right)

			switch testcase.want {
			case -1:
				if got >= 0 {
					t.Errorf("want negative value, got %d", got)
				}

			case 0:
				if got != 0 {
					t.Errorf("want 0, got %d", got)
				}

			case 1:
				if got <= 0 {
					t.Errorf("want positive value, got %d", got)
				}

			default:
				t.Errorf("Invalid direction value: %d", testcase.want)
			}
		})
	}
}
