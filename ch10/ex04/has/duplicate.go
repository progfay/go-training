package has

// Duplicate detect duplication of string slice.
// Arguments slices must be sorted.
func Duplicate(left, right []string) bool {
	var (
		ll, rl = len(left), len(right)
		li, ri = 0, 0
	)

	if ll == 0 && rl == 0 {
		return true
	}

	if ll == 0 || rl == 0 {
		return false
	}

	for {
		if ll <= li || rl <= ri {
			return false
		}

		l, r := left[li], right[ri]
		switch {
		case l == r:
			return true

		case l < r:
			li++

		case l > r:
			ri++
		}
	}
}
