package testdata_test

import (
	"testing"

	"github.com/kyuff/testdata"
	"github.com/kyuff/testdata/internal/assert"
)

func TestFillWith(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		type MyStruct struct {
			Name string
			Age  int
		}
		t.Run("should fill zero value", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillWith(t, cfg, MyStruct{})

			// assert
			assert.NotZero(t, a)
			assert.NotZero(t, a.Name)
			assert.NotZero(t, a.Age)
		})
		t.Run("should not overwrite non-zero value", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillWith(t, cfg, MyStruct{Name: "hello"})

			// assert
			assert.Equal(t, "hello", a.Name)
			assert.NotZero(t, a.Age)
		})
		t.Run("should overwrite nested pointer", func(t *testing.T) {
			// arrange
			type Nested struct {
				Name string
			}
			type MyStruct struct {
				Nested *Nested
			}
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillWith(t, cfg, MyStruct{})

			// assert
			assert.NotZero(t, a)
			assert.NotZero(t, a.Nested)
			assert.NotZero(t, a.Nested.Name)
		})
	})

	t.Run("slice", func(t *testing.T) {
		t.Run("should fill zero value", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillWith(t, cfg, []string{})

			// assert
			assert.NotEqual(t, 0, len(a))
			assert.NotZero(t, a[0])
		})
		t.Run("should not modify slice with values", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillWith(t, cfg, []string{"hello"})

			// assert
			assert.Equal(t, 1, len(a))
			assert.Equal(t, "hello", a[0])
		})
		t.Run("should fill nested slices", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillWith(t, cfg, [][][]string{})

			// assert
			assert.NotZero(t, a)
			assert.NotZero(t, a[0])
			assert.NotZero(t, a[0][0])
			assert.NotZero(t, a[0][0][0])
		})

		t.Run("pointer to slice", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillWith(t, cfg, &[]string{})

			// assert
			assert.NotNil(t, a)
			assert.NotEqual(t, 0, len(*a))
			assert.NotZero(t, (*a)[0])
		})
	})

	t.Run("map", func(t *testing.T) {
		t.Run("should fill zero value with elements", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillWith(t, cfg, map[string]string{})

			// assert
			assert.NotEqual(t, 0, len(a))
			for k, v := range a {
				assert.NotZero(t, k)
				assert.NotZero(t, v)
			}
		})
		t.Run("should not modify map with values", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillWith(t, cfg, map[string]string{"hello": "world"})

			// assert
			assert.Equal(t, "world", a["hello"])
			assert.Equal(t, 1, len(a))
		})
		t.Run("should fill nested maps", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillWith(t, cfg, map[string]map[string]string{})

			// assert
			assert.NotZero(t, a)
			for k, v := range a {
				assert.NotZero(t, k)
				assert.NotZero(t, v)
				for k2, v2 := range v {
					assert.NotZero(t, k2)
					assert.NotZero(t, v2)
				}
			}
		})
	})

	t.Run("pointer", func(t *testing.T) {
		t.Run("should fill zero value", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillWith(t, cfg, (*string)(nil))

			// assert
			assert.NotNil(t, a)
			assert.NotZero(t, *a)
		})
		t.Run("should not overwrite non-zero value", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			var st *string

			// act
			res := testdata.FillWith(t, cfg, &st)

			// assert
			assert.NotNil(t, res)
			assert.NotZero(t, *res)
		})
		t.Run("should fill pointer to a pointer", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
				val = (**string)(nil)
			)
			// act
			a := testdata.FillWith(t, cfg, val)

			// assert
			assert.NotZero(t, a)
			assert.NotZero(t, *a)
			assert.NotZero(t, **a)
		})
		t.Run("should fill pointer to a pointer to a slice", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
				val = (**[]string)(nil)
			)
			// act
			a := testdata.FillWith(t, cfg, val)

			// assert
			assert.NotZero(t, a)
			assert.NotZero(t, *a)
			assert.NotZero(t, **a)
			assert.NotEqual(t, 0, len(**a))
		})
	})
	t.Run("string", func(t *testing.T) {
		t.Run("should fill zero value", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillWith(t, cfg, "")

			// assert
			assert.NotZero(t, a)
		})
		t.Run("should not overwrite non-zero value", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillWith(t, cfg, "hello")

			// assert
			assert.Equal(t, "hello", a)
		})
		t.Run("typed string", func(t *testing.T) {
			type TypedString string
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillWith(t, cfg, TypedString(""))

			// assert
			assert.NotZero(t, a)
		})
		t.Run("should not overwrite non-zero value", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			var st *string

			// act
			res := testdata.FillWith(t, cfg, st)

			// assert
			assert.NotNil(t, res)
			assert.NotZero(t, *res)
		})
	})
}

func TestFillStickyWith(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		t.Run("should be sticky", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillStickyWith(t, cfg, "")
			b := testdata.FillWith(t, cfg, "")

			// assert
			assert.NotZero(t, a)
			assert.Equal(t, a, b)
		})
	})

	t.Run("struct", func(t *testing.T) {
		type MyStruct struct {
			Name string
			Age  int
		}
		t.Run("should be sticky", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillStickyWith(t, cfg, MyStruct{})
			b := testdata.FillWith(t, cfg, MyStruct{})

			// assert
			assert.NotZero(t, a)
			assert.Equal(t, a, b)
		})

		type AliasString string
		type StructWithTypedString struct {
			Name AliasString
		}
		t.Run("sticky string should be used in filling struct", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillStickyWith(t, cfg, AliasString(""))
			b := testdata.FillWith(t, cfg, StructWithTypedString{})

			// assert
			assert.NotZero(t, a)
			assert.NotZero(t, b.Name)
			assert.Equal(t, a, b.Name)
		})

		t.Run("should not sticky nested type when stickying struct", func(t *testing.T) {
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			// act
			a := testdata.FillStickyWith(t, cfg, StructWithTypedString{})
			b := testdata.FillWith(t, cfg, AliasString(""))

			// assert
			assert.NotZero(t, a.Name)
			assert.NotZero(t, b)
			assert.NotEqual(t, a.Name, b)
		})

	})
}
