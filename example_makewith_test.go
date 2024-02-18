package testdata_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/kyuff/testdata"
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

func ExampleMakeWith() {
	var (
		t   = &testing.T{}
		cfg = testdata.NewConfig(
			testdata.WithValues(knownDishes),
			testdata.WithRand(rand.New(rand.NewPCG(1, 2))), // stable output
		)
		_ = testdata.MakeStickyWith[City](t, cfg)
		a = testdata.MakeWith[Person](t, cfg)
		b = testdata.MakeWith[Person](t, cfg, func(person Person) Person {
			person.Age = 32
			return person
		})
	)

	fmt.Printf("A: %#v\n", a)
	fmt.Printf("B: %#v\n", b)

	// Output:
	// A: testdata_test.Person{Name:"Name-Jv4k3Dc1kHJPXsAF", Age:38380406349626614, Dish:"SPAGHETTI", City:"City-LCMNe2ur8bFrW7oM"}
	// B: testdata_test.Person{Name:"Name-0CWH6nUQnH3DqWyT", Age:32, Dish:"SPAGHETTI", City:"City-LCMNe2ur8bFrW7oM"}
}
