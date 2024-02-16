package generate

import (
	"math/rand"
	"reflect"
)

func Float64() reflect.Value {
	return reflect.ValueOf(rand.Float64())
}
