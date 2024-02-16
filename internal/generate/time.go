package generate

import (
	"reflect"
	"time"
)

func Time(typ reflect.Type) reflect.Value {
	return reflect.ValueOf(time.Now())
}
