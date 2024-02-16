package assert

import (
	"reflect"
	"regexp"
	"testing"
	"time"
)

func Equal[T comparable](t *testing.T, expected, got T) bool {
	t.Helper()
	if expected != got {
		t.Logf(`
Items was not equal
Expected: %v
     Got: %v`, expected, got)
		t.Fail()
		return false
	}
	return true
}

func NotEqual[T comparable](t *testing.T, unexpected, got T) bool {
	t.Helper()
	if unexpected == got {
		t.Logf(`
Items was equal
Expected: %v
     Got: %v`, unexpected, got)
		t.Fail()
		return false
	}
	return true
}

func NotNil(t *testing.T, got any) bool {
	t.Helper()
	if reflect.ValueOf(got).IsNil() {
		t.Logf("Expected a value, but got nil")
		t.Fail()
		return false
	}

	return true
}

func Match[T ~string](t *testing.T, expectedRE string, got T) bool {
	t.Helper()
	re, err := regexp.Compile(expectedRE)
	if err != nil {
		t.Fatalf("unexpected regexp: %s", err)
		return false
	}

	match := re.MatchString(string(got))
	if !match {
		t.Logf(`
Must match %q
       Got %q`, expectedRE, got)
		t.Fail()
		return false
	}

	return true
}

func OneOf[T comparable](t *testing.T, items []T, got T) bool {
	t.Helper()
	var found = false
	for _, item := range items {
		if item == got {
			found = true
		}
	}

	if !found {
		t.Logf("Input list: %v", items)
		t.Logf("Did not contain item: %v", got)
		t.Fail()
		return false
	}

	return true
}

func NoneZero[T any, E ~[]T](t *testing.T, got E) bool {
	t.Helper()
	for _, e := range got {
		if reflect.ValueOf(e).IsZero() {
			return false
		}
	}

	return true
}

func NotZero[T any](t *testing.T, got T) bool {
	t.Helper()
	if reflect.ValueOf(got).IsZero() {
		t.Logf("Value %T was zero: %s", got, got)
		t.Fail()
	}
	return true
}

func TimeWithinWindow(t *testing.T, expected time.Time, got time.Time, window time.Duration) bool {
	var (
		from = expected.Add(-1 * window)
		to   = expected.Add(window)
	)

	if got.Before(from) {
		t.Logf("Time was before the window by %s", from.Sub(got))
		t.Fail()
	}

	if got.After(to) {
		t.Logf("Time was after the window by %s", got.Sub(to))
		t.Fail()
	}

	return true
}

func NoError(t *testing.T, got error) bool {
	t.Helper()
	if got != nil {
		t.Logf("Unexpected error: %s", got)
		t.Fail()
		return false
	}

	return true
}

func Error(t *testing.T, got error) bool {
	t.Helper()
	if got == nil {
		t.Logf("Expected error: %s", got)
		t.Fail()
		return false
	}

	return true
}
