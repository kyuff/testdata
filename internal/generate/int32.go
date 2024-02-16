package generate

import (
	"math/rand"
	"reflect"
)

func Int32() reflect.Value {
	return reflect.ValueOf(rand.Int31())
}
