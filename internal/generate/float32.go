package generate

import (
	"math/rand"
	"reflect"
)

func Float32() reflect.Value {
	return reflect.ValueOf(rand.Float32())
}
