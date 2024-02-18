package generate

import (
	"math"
	"math/rand/v2"
	"reflect"
)

func Uint16(rand *rand.Rand) reflect.Value {
	return reflect.ValueOf(uint16(rand.UintN(math.MaxInt16)))
}
