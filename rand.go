package testdata

import "math/rand"

func randFrom[T any](values []T) T {
	if len(values) == 0 {
		var t T
		return t
	}
	return values[rand.Intn(len(values))]
}
