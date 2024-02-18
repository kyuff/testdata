package generate

import (
	"math/rand/v2"
	"reflect"
)

func Int(rand *rand.Rand) reflect.Value {
	return reflect.ValueOf(rand.Int())
}
