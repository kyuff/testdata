package generate

import (
	"math"
	"math/rand/v2"
	"reflect"
)

func Uint8(rand *rand.Rand) reflect.Value {
	return reflect.ValueOf(uint8(rand.UintN(math.MaxInt8)))
}
