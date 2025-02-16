package testdata

import (
	"reflect"
)

type testingT interface {
	Name() string
	Cleanup(fn func())
}

// Make creates a value T based on DefaultConfig
func Make[T any](t testingT, modifications ...func(d T) T) T {
	return MakeWith(t, DefaultConfig, modifications...)
}

// MakeWith creates a value T based on tge Config parameter
func MakeWith[T any](t testingT, cfg *Config, modifications ...func(d T) T) T {
	var (
		data T
	)
	data = FillWith(t, cfg, data, modifications...)
	return data
}

// MakeSticky works like Make, except values created with it, will be sticky within a t.
// That means a value of a specific type will be the same for all those types,
// even if it's a field on another Make call, or it's to a pointer to the same type.
func MakeSticky[T any](t testingT, modifications ...func(d T) T) T {
	return MakeStickyWith(t, DefaultConfig, modifications...)
}

// MakeStickyWith is similar to MakeSticky, just using cfg instead of DefaultConfig.
func MakeStickyWith[T any](t testingT, cfg *Config, modifications ...func(d T) T) T {
	var (
		typ   = reflect.TypeFor[T]()
		value = MakeWith(t, cfg, modifications...)
	)

	cfg.sticky.AddValue(t, typ, reflect.ValueOf(value))

	return value
}
