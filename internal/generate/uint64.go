package generate

import (
	"math/rand/v2"
	"reflect"
)

func Uint64() reflect.Value {
	return reflect.ValueOf(rand.Uint64())
}
