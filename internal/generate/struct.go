package generate

import "reflect"

func Struct(typ reflect.Type, maker func(typ reflect.Type) reflect.Value) reflect.Value {
	var val = reflect.Indirect(reflect.New(typ))
	for _, f := range reflect.VisibleFields(typ) {
		if !f.IsExported() {
			continue
		}
		var field = val.FieldByIndex(f.Index)
		var v = maker(field.Type())
		if v.Type().AssignableTo(field.Type()) {
			field.Set(v)
		} else {
			field.Set(v.Convert(field.Type()))
		}
	}

	return val
}
