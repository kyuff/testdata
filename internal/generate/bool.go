package generate

import (
	"math/rand/v2"
	"reflect"
)

func Bool(rand *rand.Rand) reflect.Value {
	return reflect.ValueOf(rand.Uint32()%2 == 0)
}
