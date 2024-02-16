package generate

import "reflect"

func Map(typ reflect.Type, maker func(typ reflect.Type) reflect.Value, size int) reflect.Value {
	var (
		keyType = typ.Key()
		valType = typ.Elem()
		theMap  = reflect.MakeMap(typ)
	)

	for i := 0; i < size; i++ {
		key := maker(keyType)
		if !key.Type().AssignableTo(keyType) {
			key = key.Convert(keyType)
		}

		val := maker(valType)
		if !val.Type().AssignableTo(valType) {
			val = val.Convert(valType)
		}

		theMap.SetMapIndex(key, val)
	}

	return theMap
}
