package generate

import (
	"math/rand/v2"
	"reflect"
)

func Float64(rand *rand.Rand) reflect.Value {
	return reflect.ValueOf(rand.Float64())
}
