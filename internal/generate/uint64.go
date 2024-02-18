package generate

import (
	"math/rand/v2"
	"reflect"
)

func Uint64(rand *rand.Rand) reflect.Value {
	return reflect.ValueOf(rand.Uint64())
}
