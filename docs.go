// Package testdata allows for easy creation of random data used in tests.
//
// The following features are supported:
//   - Custom types
//   - Built-in types
//   - Sticky variables
//   - Global generator modifications
//   - Local generator modifications
//
// The main entrypoint is the Make and MakeSticky functions. They are meant to be used in tests to generate
// a variable of a given type. They will use the globally defined DefaultConfig for the generation. If need be,
// it is possible to define a local Config and use it by calling MakeWith and MakeStickyWith.
//
// All Make functions accepts a list of modification functions that can be used to modify the generated value.
// This is convenient in a test, where there is a need to make it clear the input has a certain value
// that affects the expected outcome of the test.
package testdata
