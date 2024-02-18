package generate

import (
	rand2 "math/rand/v2"
	"reflect"
)

func Int64(rand *rand2.Rand) reflect.Value {
	return reflect.ValueOf(rand.Int64())
}
