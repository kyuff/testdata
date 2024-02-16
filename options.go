package testdata

import "reflect"

type Option func(cfg *Config)

// Generator will override the default value generation and instead
// use the supplied generator func for the default config.
func Generator[T any](generator func() T) {
	WithGenerator(generator)(DefaultConfig)
}

// WithGenerator will override the default value generation and instead
// use the supplied generator func for the config.
func WithGenerator[T any](generator func() T) Option {
	return func(cfg *Config) {
		var (
			t   T
			typ = reflect.TypeOf(t)
		)

		cfg.rules[typ] = func() reflect.Value {
			return reflect.ValueOf(generator())
		}
	}
}

// Values will pick one of the supplied values when
// generating a value of type T using the default config.
func Values[T any, E ~[]T](values E) {
	WithValues(values)(DefaultConfig)
}

// WithValues will pick one of the supplied values when
// generating a value of type T
func WithValues[T any, E ~[]T](values E) Option {
	return WithGenerator(func() T {
		return randFrom(values)
	})
}
