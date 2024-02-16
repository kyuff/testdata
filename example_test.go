package testdata_test

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
