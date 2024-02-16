package generate

import (
	"math"
	"math/rand/v2"
	"reflect"
)

func Uint() reflect.Value {
	return reflect.ValueOf(rand.UintN(math.MaxUint))
}
