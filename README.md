# Testdata generator for Go

[![Build Status](https://github.com/kyuff/testdata/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/kyuff/testdata/actions/workflows/go.yml)
[![Report Card](https://goreportcard.com/badge/github.com/kyuff/testdata)](https://goreportcard.com/report/github.com/kyuff/testdata/)
[![Go Reference](https://pkg.go.dev/badge/github.com/kyuff/testdata.svg)](https://pkg.go.dev/github.com/kyuff/testdata)
[![codecov](https://codecov.io/gh/kyuff/testdata/graph/badge.svg?token=GA4GSQCLZE)](https://codecov.io/gh/kyuff/testdata)

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
called each time there is a need to generate that specific type. This is a method to override the default generator.

## Example

````go
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
````

Find more detailed examples [here](example_make_test.go) and [here](example_makewith_test.go).
