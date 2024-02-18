package testdata

import (
	"math/rand/v2"
	"reflect"
	"time"

	"github.com/kyuff/testdata/internal/generate"
	"github.com/kyuff/testdata/internal/sticky"
)

// DefaultConfig is the Config that will be used with Make and MakeSticky.
// Configure it using the options that are not prefixed using With, ie Values or Generator.
var DefaultConfig *Config

// init the DefaultConfig
func init() {
	DefaultConfig = NewConfig()
}

// NewConfig constructs a Config that can be used with MakeWith and MakeStickyWith.
// Configure it using one ore more Option that is prefixed using With*.
func NewConfig(opts ...Option) *Config {
	cfg := &Config{
		rules:  make(map[reflect.Type]func(r *rand.Rand) reflect.Value),
		sticky: sticky.New(),
		rand:   rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())),
	}
	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

// Config for testdata. Use either the DefaultConfig or create one with NewConfig.
type Config struct {
	rules  map[reflect.Type]func(r *rand.Rand) reflect.Value
	sticky *sticky.Manager
	rand   *rand.Rand
}

func (cfg *Config) make(t testingT, typ reflect.Type) reflect.Value {
	stickyValue, isSticky := cfg.sticky.HasValue(t, typ)
	if isSticky {
		return stickyValue
	}

	rule, ok := cfg.rules[typ]
	if ok {
		return rule(cfg.rand)
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
		return generate.Time(cfg.rand, typ)
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
		return generate.String(cfg.rand, typ, 16)
	case reflect.Int:
		return generate.Int(cfg.rand)
	case reflect.Int8:
		return generate.Int8(cfg.rand)
	case reflect.Int16:
		return generate.Int16(cfg.rand)
	case reflect.Int32:
		return generate.Int32(cfg.rand)
	case reflect.Int64:
		return generate.Int64(cfg.rand)
	case reflect.Bool:
		return generate.Bool(cfg.rand)
	case reflect.Uint:
		return generate.Uint(cfg.rand)
	case reflect.Uint8:
		return generate.Uint8(cfg.rand)
	case reflect.Uint16:
		return generate.Uint16(cfg.rand)
	case reflect.Uint32:
		return generate.Uint32(cfg.rand)
	case reflect.Uint64:
		return generate.Uint64(cfg.rand)
	case reflect.Float32:
		return generate.Float32(cfg.rand)
	case reflect.Float64:
		return generate.Float64(cfg.rand)
	default:
		return reflect.Zero(typ)
	}
}
