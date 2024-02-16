package generate

import (
	"math"
	"math/rand"
	"reflect"
)

func Int8() reflect.Value {
	return reflect.ValueOf(int8(rand.Intn(math.MaxInt8)))
}
