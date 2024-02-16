package generate

import (
	"math/rand/v2"
	"reflect"
)

func Uint32() reflect.Value {
	return reflect.ValueOf(rand.Uint32())
}
