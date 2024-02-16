package generate

import (
	"math"
	"math/rand"
	"reflect"
)

func Int16() reflect.Value {
	return reflect.ValueOf(int16(rand.Intn(math.MaxInt16)))
}
