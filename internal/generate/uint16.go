package generate

import (
	"math"
	"math/rand/v2"
	"reflect"
)

func Uint16() reflect.Value {
	return reflect.ValueOf(uint16(rand.UintN(math.MaxInt16)))
}
