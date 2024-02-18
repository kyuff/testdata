package generate

import (
	"math/rand/v2"
	"reflect"
)

func Uint32(rand *rand.Rand) reflect.Value {
	return reflect.ValueOf(rand.Uint32())
}
