package testdata_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/kyuff/testdata"
)

func ExampleMakeWith() {
	var (
		t   = &testing.T{}
		cfg = testdata.NewConfig(
			testdata.WithRand(rand.New(rand.NewPCG(1, 2))), // stable output
			testdata.WithValues(knownDishes),
			testdata.WithGenerator(func(rand *rand.Rand) Name {
				return Name(fmt.Sprintf("My own name: %d", rand.IntN(100)))
			}),
			testdata.WithGenerator(func(rand *rand.Rand) Age {
				return Age(rand.IntN(100))
			}),
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
	// A: testdata_test.Person{Name:"My own name: 73", Age:50, Dish:"SPAGHETTI", City:"City-LCMNe2ur8bFrW7oM", Note:"string-k3Dc1kHJPXsAFv0C"}
	// B: testdata_test.Person{Name:"My own name: 94", Age:32, Dish:"SPAGHETTI", City:"City-LCMNe2ur8bFrW7oM", Note:"string-nUQnH3DqWyTPPTEi"}
}
