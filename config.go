package testdata

import (
	"reflect"
	"time"

	"github.com/kyuff/testdata/internal/generate"
	"github.com/kyuff/testdata/internal/sticky"
)

var DefaultConfig *Config

func init() {
	DefaultConfig = NewConfig()
}

func NewConfig(opts ...Option) *Config {
	cfg := &Config{
		rules:  make(map[reflect.Type]func() reflect.Value),
		sticky: sticky.New(),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

type Config struct {
	rules  map[reflect.Type]func() reflect.Value
	sticky *sticky.Manager
}

func (cfg *Config) make(t testingT, typ reflect.Type) reflect.Value {
	stickyValue, isSticky := cfg.sticky.HasValue(t, typ)
	if isSticky {
		return stickyValue
	}

	rule, ok := cfg.rules[typ]
	if ok {
		return rule()
	}

	var pointer = typ.Kind() == reflect.Pointer
	if pointer {
		typ = typ.Elem()
	}

	var v = cfg.generateBuiltIn(t, typ)
	if pointer {
		return generate.Pointer(v)
	}

	return v
}

var timeType = reflect.TypeOf(time.Time{})

func (cfg *Config) generateBuiltIn(t testingT, typ reflect.Type) reflect.Value {
	if timeType.ConvertibleTo(typ) {
		return generate.Time(typ)
	}
	var maker = func(typ reflect.Type) reflect.Value {
		return cfg.make(t, typ)
	}
	switch typ.Kind() {
	case reflect.Struct:
		return generate.Struct(typ, maker)
	case reflect.Slice:
		return generate.Slice(typ, maker, 5)
	case reflect.Map:
		return generate.Map(typ, maker, 5)
	case reflect.String:
		return generate.String(typ, 16)
	case reflect.Int:
		return generate.Int()
	case reflect.Int8:
		return generate.Int8()
	case reflect.Int16:
		return generate.Int16()
	case reflect.Int32:
		return generate.Int32()
	case reflect.Int64:
		return generate.Int64()
	case reflect.Bool:
		return generate.Bool()
	case reflect.Uint:
		return generate.Uint()
	case reflect.Uint8:
		return generate.Uint8()
	case reflect.Uint16:
		return generate.Uint16()
	case reflect.Uint32:
		return generate.Uint32()
	case reflect.Uint64:
		return generate.Uint64()
	case reflect.Float32:
		return generate.Float32()
	case reflect.Float64:
		return generate.Float64()
	default:
		return reflect.Zero(typ)
	}
}
