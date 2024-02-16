package generate

import (
	"math/rand/v2"
	"reflect"
)

func Bool() reflect.Value {
	return reflect.ValueOf(rand.Uint32()%2 == 0)
}
