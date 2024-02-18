package generate

import (
	"math"
	"math/rand/v2"
	"reflect"
)

func Int16(rand *rand.Rand) reflect.Value {
	return reflect.ValueOf(int16(rand.IntN(math.MaxInt16)))
}
