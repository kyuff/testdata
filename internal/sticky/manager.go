package sticky

import (
	"reflect"
	"sync"

	"github.com/kyuff/testdata/internal/generate"
)

type testingT interface {
	Name() string
	Cleanup(fn func())
}

type (
	TestValues map[reflect.Type]reflect.Value
	TestScope  map[string]TestValues
)

func New() *Manager {
	return &Manager{
		scopes: make(TestScope),
	}
}

type Manager struct {
	mu     sync.RWMutex
	scopes TestScope
}

func (mgr *Manager) HasValue(t testingT, typ reflect.Type) (reflect.Value, bool) {
	mgr.mu.RLock()
	defer mgr.mu.RUnlock()

	var pointer = typ.Kind() == reflect.Pointer
	if pointer {
		typ = typ.Elem()
	}

	values, ok := mgr.scopes[t.Name()]
	if !ok {
		return reflect.ValueOf(nil), false
	}

	value, ok := values[typ]
	if !ok {
		return reflect.ValueOf(nil), false
	}

	if pointer {
		return generate.Pointer(value), true
	}

	return value, true
}

func (mgr *Manager) cleanup(scope string) func() {
	return func() {
		mgr.mu.Lock()
		defer mgr.mu.Unlock()
		delete(mgr.scopes, scope)
	}
}
func (mgr *Manager) AddValue(t testingT, typ reflect.Type, val reflect.Value) {
	mgr.mu.Lock()
	defer mgr.mu.Unlock()

	scope, ok := mgr.scopes[t.Name()]
	if !ok {
		scope = make(TestValues)
		t.Cleanup(mgr.cleanup(t.Name()))
	}

	scope[typ] = val
	mgr.scopes[t.Name()] = scope
}
