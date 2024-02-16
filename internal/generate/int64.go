package generate

import (
	"math/rand"
	"reflect"
)

func Int64() reflect.Value {
	return reflect.ValueOf(rand.Int63())
}
