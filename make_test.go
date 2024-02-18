package testdata_test

import (
	"math/rand/v2"
	"testing"
	"time"

	"github.com/kyuff/testdata"
	"github.com/kyuff/testdata/internal/assert"
)

func TestMake(t *testing.T) {
	t.Parallel()

	t.Run("Sticky", func(t *testing.T) {
		t.Parallel()
		type ID string
		type State string
		type Person struct {
			ID    ID
			State State
		}

		t.Run("prevent share between tests", func(t *testing.T) {
			var (
				cfg     = testdata.NewConfig()
				results = make(chan ID, 2)
			)

			t.Run("a", func(t *testing.T) {
				var (
					id = testdata.MakeStickyWith[ID](t, cfg)
				)

				// act
				person := testdata.MakeWith[Person](t, cfg)

				// assert
				assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", id)
				assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", person.ID)
				assert.Match(t, "^State-[a-zA-Z0-9]{16}$", person.State)
				if assert.Equal(t, id, person.ID) {
					results <- id
				}
			})

			t.Run("b", func(t *testing.T) {
				var (
					id = testdata.MakeStickyWith[ID](t, cfg)
				)

				// act
				person := testdata.MakeWith[Person](t, cfg)

				// assert
				assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", id)
				assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", person.ID)
				assert.Match(t, "^State-[a-zA-Z0-9]{16}$", person.State)
				if assert.Equal(t, id, person.ID) {
					results <- id
				}
			})

			r1 := <-results
			r2 := <-results
			assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", r1)
			assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", r2)
			assert.NotEqual(t, r1, r2)
		})

		t.Run("sticky pointer value", func(t *testing.T) {
			var (
				cfg = testdata.NewConfig()
				id  = testdata.MakeStickyWith[ID](t, cfg)
			)

			// act
			got := testdata.MakeWith[*ID](t, cfg)

			// assert
			if assert.NotNil(t, got) {
				assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", id)
				assert.Equal(t, id, *got)
			}
		})

		t.Run("two calls", func(t *testing.T) {
			var (
				cfg = testdata.NewConfig()
				id  = testdata.MakeStickyWith[ID](t, cfg)
			)

			// act
			got := testdata.MakeStickyWith[ID](t, cfg)

			// assert
			assert.Equal(t, id, got)
		})

		t.Run("another sticky", func(t *testing.T) {
			var (
				cfg = testdata.NewConfig()
				id  = testdata.MakeStickyWith[ID](t, cfg)
			)

			// act
			got := testdata.MakeStickyWith[Person](t, cfg)

			// assert
			assert.Equal(t, id, got.ID)
		})

		t.Run("with modifications", func(t *testing.T) {
			var (
				cfg = testdata.NewConfig()
				id  = testdata.MakeStickyWith[ID](t, cfg, func(d ID) ID {
					return "my id"
				})
			)

			// act
			got := testdata.MakeWith[Person](t, cfg)

			// assert
			assert.Equal(t, "my id", id)
			assert.Equal(t, "my id", got.ID)
		})

		t.Run("with modifications on make", func(t *testing.T) {
			var (
				cfg = testdata.NewConfig()
				id  = testdata.MakeStickyWith[ID](t, cfg)
			)

			// act
			got := testdata.MakeWith[Person](t, cfg, func(d Person) Person {
				d.State = "my state"
				return d
			})

			// assert
			assert.Equal(t, id, got.ID)
			assert.Equal(t, "my state", got.State)
		})

		t.Run("with sticky generator", func(t *testing.T) {
			var (
				cfg = testdata.NewConfig(testdata.WithGenerator(func(r *rand.Rand) ID {
					return "my id"
				}))
				id = testdata.MakeStickyWith[ID](t, cfg)
			)

			// act
			got := testdata.MakeWith[Person](t, cfg)

			// assert
			assert.Equal(t, "my id", id)
			assert.Equal(t, "my id", got.ID)
		})
	})

	t.Run("Options", func(t *testing.T) {
		t.Parallel()
		t.Run("Values", func(t *testing.T) {
			t.Parallel()
			// arrange
			type TypedString string
			var (
				knownTypedString = []TypedString{"A", "B", "C", "D"}
				cfg              = testdata.NewConfig(
					testdata.WithValues(knownTypedString),
				)
			)

			// act
			got := testdata.MakeWith[TypedString](t, cfg)

			// assert
			assert.OneOf(t, knownTypedString, got)
		})

		t.Run("Values empty", func(t *testing.T) {
			t.Parallel()
			// arrange
			type TypedString string
			var (
				knownTypedString []TypedString
				emptyTypedString TypedString
				cfg              = testdata.NewConfig(
					testdata.WithValues(knownTypedString),
				)
			)

			// act
			got := testdata.MakeWith[TypedString](t, cfg)

			// assert
			assert.Equal(t, emptyTypedString, got)
		})

		t.Run("Generator", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				expected = rand.Int64()
				cfg      = testdata.NewConfig(
					testdata.WithGenerator(func(r *rand.Rand) int64 {
						return expected
					}),
				)
			)

			// act
			got := testdata.MakeWith[int64](t, cfg)

			// assert
			assert.Equal(t, expected, got)
		})

		t.Run("Rand", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig(
					testdata.WithRand(rand.New(rand.NewPCG(1, 2))),
				)
			)

			// act
			got := testdata.MakeWith[int](t, cfg)

			// assert
			assert.Equal(t, 4969059760275911952, got)
		})
	})

	t.Run("Modifications", func(t *testing.T) {
		t.Parallel()
		type ID string
		type Country string
		type Data struct {
			ID      ID
			Age     int
			Country Country
		}

		t.Run("apply", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			// act
			got := testdata.MakeWith(t, cfg, func(d Data) Data {
				d.ID = "My ID"
				d.Age = 5
				return d
			})

			// assert
			assert.Equal(t, "My ID", got.ID)
			assert.Equal(t, 5, got.Age)
			assert.Match(t, "^Country-[a-zA-Z0-9]{16}$", got.Country)
		})

		t.Run("apply pointer", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			// act
			got := testdata.MakeWith(t, cfg, func(d *Data) *Data {
				d.ID = "My ID"
				d.Age = 5
				return d
			})

			// assert
			if assert.NotNil(t, got) {
				assert.Equal(t, "My ID", got.ID)
				assert.Equal(t, 5, got.Age)
				assert.Match(t, "^Country-[a-zA-Z0-9]{16}$", got.Country)
			}
		})
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		t.Run("make", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			// act
			got := testdata.MakeWith[string](t, cfg)

			// assert
			assert.Match(t, "^string-[a-zA-Z0-9]{16}$", got)
		})

		t.Run("typed", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			type TypedString string

			// act
			got := testdata.MakeWith[TypedString](t, cfg)

			// assert
			assert.Match(t, "^TypedString-[a-zA-Z0-9]{16}$", got)
		})

		t.Run("typed pointer", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			type TypedString string

			// act
			got := testdata.MakeWith[*TypedString](t, cfg)

			// assert
			if assert.NotNil(t, got) {
				assert.Match(t, "^TypedString-[a-zA-Z0-9]{16}$", *got)
			}
		})

		t.Run("string pointer", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			// act
			got := testdata.MakeWith[*string](t, cfg)

			// assert
			if assert.NotNil(t, got) {
				assert.Match(t, "^string-[a-zA-Z0-9]{16}$", *got)
			}
		})
	})

	t.Run("struct", func(t *testing.T) {
		t.Parallel()
		type ID string
		type Name string
		type Person struct {
			ID    ID
			Age   int
			Names []Name
		}

		t.Run("make", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			// act
			got := testdata.MakeWith[Person](t, cfg)

			// assert
			assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", got.ID)
			if assert.NotNil(t, got.Names) {
				assert.Equal(t, 5, len(got.Names))
			}
		})

		t.Run("make pointer", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			// act
			got := testdata.MakeWith[*Person](t, cfg)

			// assert
			if assert.NotNil(t, got) {
				assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", got.ID)
				if assert.NotNil(t, got.Names) {
					assert.Equal(t, 5, len(got.Names))
				}
			}
		})

		t.Run("ignore private field", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			type Agent struct {
				ID       ID
				CodeName Name
				realName Name
			}

			// act
			got := testdata.MakeWith[Agent](t, cfg)

			// assert
			assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", got.ID)
			assert.Match(t, "^Name-[a-zA-Z0-9]{16}$", got.CodeName)
			assert.Match(t, "^$", got.realName)
		})

		t.Run("nested", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			type Family struct {
				ParentA Person
				ParentB Person

				StepParentA *Person
				StepParentB *Person
			}

			// act
			got := testdata.MakeWith[Family](t, cfg)

			// assert
			assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", got.ParentA.ID)
			assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", got.ParentB.ID)
			if assert.NotNil(t, got.StepParentA) {
				assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", got.StepParentA.ID)
			}
			if assert.NotNil(t, got.StepParentB) {
				assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", got.StepParentB.ID)
			}
		})
	})

	t.Run("slice", func(t *testing.T) {
		t.Parallel()
		t.Run("make", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			// act
			got := testdata.MakeWith[[]string](t, cfg)

			// assert
			assert.Equal(t, 5, len(got))
			assert.NoneZero(t, got)
		})

		t.Run("make typed element", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			type Age int

			// act
			got := testdata.MakeWith[[]Age](t, cfg)

			// assert
			if assert.NotNil(t, got) {
				assert.Equal(t, 5, len(got))
				assert.NoneZero(t, got)
			}
		})

		t.Run("make typed slice", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			type Age int
			type Ages []Age
			// act
			got := testdata.MakeWith[Ages](t, cfg)

			// assert
			if assert.NotNil(t, got) {
				assert.Equal(t, 5, len(got))
				assert.NoneZero(t, got)
			}
		})

		t.Run("make typed pointer element", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			type Age int

			// act
			got := testdata.MakeWith[[]*Age](t, cfg)

			// assert
			if assert.NotNil(t, got) {
				assert.Equal(t, 5, len(got))
				assert.NoneZero(t, got)
			}
		})

		t.Run("make struct elements", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			type ID string
			type Data struct {
				ID  ID
				Age int
			}

			// act
			got := testdata.MakeWith[[]Data](t, cfg)

			// assert
			if assert.NotNil(t, got) {
				assert.Equal(t, 5, len(got))
				assert.NoneZero(t, got)
			}
		})

		t.Run("make pointer", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			type Age int

			// act
			got := testdata.MakeWith[*[]Age](t, cfg)

			// assert
			if assert.NotNil(t, got) {
				assert.Equal(t, 5, len(*got))
			}
		})

	})

	t.Run("map", func(t *testing.T) {
		t.Parallel()
		t.Run("make", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			// act
			got := testdata.MakeWith[map[string]int](t, cfg)

			// assert
			if assert.NotNil(t, got) {
				assert.Equal(t, 5, len(got))
			}
		})

		t.Run("make typed key values", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			type ID string
			type Name string
			type Person struct {
				Age  int
				Name Name
			}

			// act
			got := testdata.MakeWith[map[ID]Person](t, cfg)

			// assert
			if assert.NotNil(t, got) && assert.Equal(t, 5, len(got)) {
				for key, value := range got {
					assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", key)
					assert.Match(t, "^Name-[a-zA-Z0-9]{16}$", value.Name)
				}
			}
		})

		t.Run("make typed map", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)
			type ID string
			type Name string
			type Person struct {
				Age  int
				Name Name
			}
			type Register map[ID]Person

			// act
			got := testdata.MakeWith[Register](t, cfg)

			// assert
			if assert.NotNil(t, got) && assert.Equal(t, 5, len(got)) {
				for key, value := range got {
					assert.Match(t, "^ID-[a-zA-Z0-9]{16}$", key)
					assert.Match(t, "^Name-[a-zA-Z0-9]{16}$", value.Name)
				}
			}
		})

		t.Run("make pointer", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			// act
			got := testdata.MakeWith[*map[string]int](t, cfg)

			// assert
			if assert.NotNil(t, got) {
				assert.Equal(t, 5, len(*got))
			}
		})

		t.Run("make pointer values", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			// act
			got := testdata.MakeWith[map[string]*int](t, cfg)

			// assert
			if assert.NotNil(t, got) && assert.Equal(t, 5, len(got)) {
				for key, value := range got {
					assert.Match(t, "^string-[a-zA-Z0-9]{16}$", key)
					assert.NotNil(t, value)
				}
			}
		})
	})

	t.Run("time.Time", func(t *testing.T) {
		t.Parallel()
		const fourYears = time.Hour * 24 * 360 * 4
		t.Run("make", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			// act
			got := testdata.MakeWith[time.Time](t, cfg)

			// assert
			assert.TimeWithinWindow(t, time.Now(), got, fourYears)
		})

		t.Run("make pointer", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			// act
			got := testdata.MakeWith[*time.Time](t, cfg)

			// assert
			if assert.NotNil(t, got) {
				assert.TimeWithinWindow(t, time.Now(), *got, fourYears)
			}
		})

		t.Run("make typed", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			type TestTime time.Time

			// act
			got := testdata.MakeWith[TestTime](t, cfg)

			// assert
			assert.TimeWithinWindow(t, time.Now(), time.Time(got), fourYears)
		})

	})

	t.Run("simple/types", func(t *testing.T) {
		t.Run("make", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			type Types struct {
				Bool    bool
				Int     int
				Int8    int8
				Int16   int16
				Int32   int32
				Int64   int64
				Uint    uint
				Uint8   uint8
				Uint16  uint16
				Uint32  uint32
				Uint64  uint64
				Float32 float32
				Float64 float64
				String  string
			}

			// act
			got := testdata.MakeWith[Types](t, cfg)

			// assert
			assert.NotZero(t, got.Int)
			assert.NotZero(t, got.Int8)
			assert.NotZero(t, got.Int16)
			assert.NotZero(t, got.Int32)
			assert.NotZero(t, got.Int64)
			assert.NotZero(t, got.Uint)
			assert.NotZero(t, got.Uint8)
			assert.NotZero(t, got.Uint16)
			assert.NotZero(t, got.Uint32)
			assert.NotZero(t, got.Uint64)
			assert.NotZero(t, got.Float32)
			assert.NotZero(t, got.Float64)
			assert.NotZero(t, got.String)
		})

		t.Run("typed", func(t *testing.T) {
			t.Parallel()
			// arrange
			var (
				cfg = testdata.NewConfig()
			)

			type Bool bool
			type Int int
			type Int8 int8
			type Int16 int16
			type Int32 int32
			type Int64 int64
			type Uint uint
			type Uint8 uint8
			type Uint16 uint16
			type Uint32 uint32
			type Uint64 uint64
			type Float32 float32
			type Float64 float64
			type String string

			type Types struct {
				Bool    Bool
				Int     Int
				Int8    Int8
				Int16   Int16
				Int32   Int32
				Int64   Int64
				Uint    Uint
				Uint8   Uint8
				Uint16  Uint16
				Uint32  Uint32
				Uint64  Uint64
				Float32 Float32
				Float64 Float64
				String  String
			}

			// act
			got := testdata.MakeWith[Types](t, cfg)

			// assert
			assert.NotZero(t, got.Int)
			assert.NotZero(t, got.Int8)
			assert.NotZero(t, got.Int16)
			assert.NotZero(t, got.Int32)
			assert.NotZero(t, got.Int64)
			assert.NotZero(t, got.Uint)
			assert.NotZero(t, got.Uint8)
			assert.NotZero(t, got.Uint16)
			assert.NotZero(t, got.Uint32)
			assert.NotZero(t, got.Uint64)
			assert.NotZero(t, got.Float32)
			assert.NotZero(t, got.Float64)
			assert.NotZero(t, got.String)
		})
	})

}
