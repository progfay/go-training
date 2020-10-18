package min

import "math"

func Min(values ...int64) int64 {
	var min int64 = math.MaxInt64
	for _, v := range values {
		if min > v {
			min = v
		}
	}
	return min
}

func MinRequiresAtLeastOneArg(value int64, values ...int64) int64 {
	var min int64 = value
	for _, v := range values {
		if min > v {
			min = v
		}
	}
	return min
}
