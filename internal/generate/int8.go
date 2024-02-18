package generate

import (
	"math"
	"math/rand/v2"
	"reflect"
)

func Int8(rand *rand.Rand) reflect.Value {
	return reflect.ValueOf(int8(rand.IntN(math.MaxInt8)))
}
