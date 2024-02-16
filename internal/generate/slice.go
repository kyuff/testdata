package generate

import "reflect"

func Slice(typ reflect.Type, maker func(typ reflect.Type) reflect.Value, size int) reflect.Value {
	var (
		eleType  = typ.Elem()
		theSlice = reflect.New(reflect.SliceOf(eleType)).Elem()
	)

	for i := 0; i < size; i++ {
		v := maker(eleType)
		if !v.Type().AssignableTo(eleType) {
			v = v.Convert(eleType)
		}

		theSlice = reflect.Append(theSlice, v)
	}

	return theSlice
}
