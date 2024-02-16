package generate

import (
	"math/rand"
	"reflect"
)

func Int() reflect.Value {
	return reflect.ValueOf(rand.Int())
}
