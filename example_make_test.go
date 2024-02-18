package testdata_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/kyuff/testdata"
)

// init the testdata genration
func init() {
	testdata.Rand(rand.New(rand.NewPCG(1, 2))) // stable output
	testdata.Values(knownDishes)
	testdata.Generator(func(rand *rand.Rand) Name {
		return Name(fmt.Sprintf("My own name: %d", rand.IntN(100)))
	})
	testdata.Generator(func(rand *rand.Rand) Age {
		return Age(rand.IntN(100))
	})
}

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
	Note string
}

func ExampleMake() {
	var (
		t = &testing.T{}
		_ = testdata.MakeSticky[City](t)
		a = testdata.Make[Person](t)
		b = testdata.Make[Person](t, func(person Person) Person {
			person.Age = 32
			return person
		})
	)

	fmt.Printf("A: %#v\n", a)
	fmt.Printf("B: %#v\n", b)

	// Output:
	// A: testdata_test.Person{Name:"My own name: 73", Age:50, Dish:"SPAGHETTI", City:"City-LCMNe2ur8bFrW7oM", Note:"string-k3Dc1kHJPXsAFv0C"}
	// B: testdata_test.Person{Name:"My own name: 94", Age:32, Dish:"SPAGHETTI", City:"City-LCMNe2ur8bFrW7oM", Note:"string-nUQnH3DqWyTPPTEi"}
}
