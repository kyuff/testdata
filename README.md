# Testdata generator for Go

Library to generate random and typed data for tests.

## Use case

In some types of tests it is convenient to randomize input data in order to avoid a hidden reliance on specific values.
Especially in code bases that model domain specific values using named types, this can be cumbersome to do by hand.

This library provides a simple and convenient way to configure value types and generate them for specific tests.

## FAQ

### Can I reuse a `testdata.Config` between tests?

Yes! There is even a convenient `testdata.DefaultConfig` you can use.
You can configure the default config using the convenient testdata.Values functions. Alternatively you can use options
that is prefixed using `With` when setting up your own `Config`.

### Is it ok to use `testdata.Make()`?

Yes, the `testdata.MakeWith()` is only meant if you don't want to use the default config.

### Can I control how values are generated?

Yes, use the option `testdata.WithGenerator` which accepts a func that provides a specific type. This func will be
called each type there is a need to generate that specific type. This is a method to override the default generator.

## Example

````go
import (
	"fmt"
	"testing"

	"github.com/kyuff/testdata"
	"github.com/kyuff/testdata/internal/assert"
)

type Name string
type Age int
type City string
type Dish string

const (
	Spaghetti   Dish = "SPAGHETTI"
	CaesarSalad Dish = "CAESAR_SALAD"
	Milkshake   Dish = "MILKSHAKE"
)

var knownDishes = []Dish{Spaghetti, CaesarSalad, Milkshake}

type Person struct {
	Name Name
	Age  Age
	Dish Dish
	City City
}

func TestExample(t *testing.T) {
	// arrange
	var (
		cfg = testdata.NewConfig(
			testdata.WithValues(knownDishes),
		)
		_ = testdata.MakeStickyWith[City](t, cfg)
		a = testdata.MakeWith[Person](t, cfg)
		b = testdata.MakeWith[Person](t, cfg)
	)

	// act
	funcUnderTest(a, b)

	// assert
	assert.Equal(t, a.City, b.City)
}

func funcUnderTest(a, b Person) {
	fmt.Printf("A: %#v\n", a)
	fmt.Printf("B: %#v\n", b)
}
````

Output:

A: testdata_test.Person{Name:"Name-K1t8IV7mGP0x5GZ8", Age:9127679439778925705, Dish:"CAESAR_SALAD", City:"City-K7146TOeLlK8Vz3g"}

B: testdata_test.Person{Name:"Name-ccwupn3InlxOzDTT", Age:6402903903958643118, Dish:"MILKSHAKE", City:"City-K7146TOeLlK8Vz3g"}

