package max

import "math"

func Max(values ...int64) int64 {
	var max int64 = math.MinInt64
	for _, v := range values {
		if max < v {
			max = v
		}
	}
	return max
}

func MaxRequiresAtLeastOneArg(value int64, values ...int64) int64 {
	var max int64 = value
	for _, v := range values {
		if max < v {
			max = v
		}
	}
	return max
}
