package generate

import (
	"math/rand/v2"
	"reflect"
)

func Float32(rand *rand.Rand) reflect.Value {
	return reflect.ValueOf(rand.Float32())
}
