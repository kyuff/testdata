package generate

import (
	"math"
	"math/rand/v2"
	"reflect"
)

func Uint(rand *rand.Rand) reflect.Value {
	return reflect.ValueOf(rand.UintN(math.MaxUint))
}
