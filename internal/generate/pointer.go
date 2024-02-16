package generate

import "reflect"

func Pointer(val reflect.Value) reflect.Value {
	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Data",
			Type: val.Type(),
		},
	})

	v := reflect.New(typ).Elem()
	v.Field(0).Set(val)
	return v.Field(0).Addr()
}
