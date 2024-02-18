package testdata

import (
	"math/rand/v2"
	"reflect"
)

// Option to customize a Config.
type Option func(cfg *Config)

// Generator will override the default value generation and instead
// use the supplied generator func for DefaultConfig.
func Generator[T any](generator func(r *rand.Rand) T) {
	WithGenerator(generator)(DefaultConfig)
}

// WithGenerator will override the default value generation and instead
// use the supplied generator func for the Config.
func WithGenerator[T any](generator func(r *rand.Rand) T) Option {
	return func(cfg *Config) {
		var (
			t   T
			typ = reflect.TypeOf(t)
		)

		cfg.rules[typ] = func(r *rand.Rand) reflect.Value {
			return reflect.ValueOf(generator(r))
		}
	}
}

// Values will pick one of the supplied values when
// generating a value of type T using DefaultConfig.
func Values[T any, E ~[]T](values E) {
	WithValues(values)(DefaultConfig)
}

// WithValues will pick one of the supplied values when
// generating a value of type T
func WithValues[T any, E ~[]T](values E) Option {
	return WithGenerator(func(r *rand.Rand) T {
		return randFrom(r, values)
	})
}

// Rand will use the provided *rand.Rand when generating
// testdata using DefaultConfig.
func Rand(r *rand.Rand) {
	WithRand(r)(DefaultConfig)
}

// WithRand will use the provided *rand.Rand when generating
// testdata using the constructed Config.
func WithRand(r *rand.Rand) Option {
	return func(cfg *Config) {
		cfg.rand = r
	}
}
