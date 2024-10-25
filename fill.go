package testdata

import (
	"reflect"
)

// FillSticky works like Fill, except values created with it, will be sticky within a t.
// That means a value of a specific type will be the same for all those types,
// even if it's a field on another Fill call, or it's to a pointer to the same type.
func FillSticky[T any](t testingT, in T, modifications ...func(d T) T) T {
	return FillStickyWith(t, DefaultConfig, in, modifications...)
}

// FillStickyWith is similar to FillSticky, just using cfg instead of DefaultConfig.
func FillStickyWith[T any](t testingT, cfg *Config, in T, modifications ...func(d T) T) T {
	var (
		typ   = reflect.TypeFor[T]()
		value = FillWith(t, cfg, in, modifications...)
	)

	cfg.sticky.AddValue(t, typ, reflect.ValueOf(value))

	return value
}

// Fill returns a new value where zero-valued fields / value from the in parameter are filled with random values
// Pointers that are nil will be set to a pointer that references a non zero value. This is also true for pointers to pointers
// If initialValue is a pointer to a zero-value, Fill will return a pointer to a non-zero value
// Slice and Map types that are empty will be filled with values
// Instead, it will return a new value where zero-valued fields / value are filled with random values
func Fill[T any](t testingT, in T, modifications ...func(d T) T) T {
	return FillWith(t, DefaultConfig, in, modifications...)
}

// Fill is similar to Fill, just using cfg instead of DefaultConfig.
func FillWith[T any](t testingT, cfg *Config, initialValue T, modifications ...func(d T) T) T {
	var (
		value = initialValue
		val   = reflect.ValueOf(&value)
	)
	fillZeroValues(t, cfg, val)

	for _, modify := range modifications {
		value = modify(value)
	}

	return value
}

func convert(a reflect.Value, b reflect.Value) reflect.Value {
	if b.Type().ConvertibleTo(a.Type()) || b.CanConvert(a.Type()) {
		return b.Convert(a.Type())
	}
	return b
}

func fillZeroValues(t testingT, cfg *Config, original reflect.Value) {
	switch original.Kind() {
	case reflect.Struct:
		if original.IsZero() {
			newValue := cfg.make(t, original.Type())
			original.Set(convert(original, newValue))
		}
		for _, field := range reflect.VisibleFields(original.Type()) {
			if !field.IsExported() {
				//ignore fields that are not exported
				continue
			}

			valueField := original.FieldByIndex(field.Index)
			if valueField.IsZero() {
				fillZeroValues(t, cfg, valueField)
			}
		}
	case reflect.Ptr:
		if original.IsNil() {
			newValue := cfg.make(t, original.Type())
			fillZeroValues(t, cfg, newValue)
			original.Set(convert(original, newValue))
		} else {
			fillZeroValues(t, cfg, original.Elem())
		}
	case reflect.Slice, reflect.Map:
		if original.IsNil() {
			newValue := cfg.make(t, original.Type())
			original.Set(convert(original, newValue))
		}
		if original.Len() == 0 {
			newValue := cfg.make(t, original.Type())
			original.Set(convert(original, newValue))
		}
	default:
		if original.IsZero() {
			newValue := cfg.make(t, original.Type())
			original.Set(convert(original, newValue))
		}
	}
}
