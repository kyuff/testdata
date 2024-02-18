package testdata

import "math/rand/v2"

func randFrom[T any](r *rand.Rand, values []T) T {
	if len(values) == 0 {
		var t T
		return t
	}
	return values[r.IntN(len(values))]
}
