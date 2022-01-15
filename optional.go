package goptional

import (
	"errors"
	"fmt"
)

// EmptyOptional is the error returned when Unwrap is called on an Optional without a value.
var EmptyOptional = errors.New("optional is empty")

// Optional is a monad that wraps a value to either be present or empty.
// This can be used as a way to represent data that may or may not exist, or as returns values that may be optional.
type Optional[T any] struct {
	present bool // TODO(eth-p): If specialization is ever added, specialize to use nil value as empty.
	value   T
}

// String prints the string representation of the Optional's value.
// This will be whatever fmt.Sprintf("%v") prints if the Optional is Present, or "(None)" if Empty.
func (o Optional[T]) String() string {
	if o.present {
		return fmt.Sprintf("%v", o.value)
	} else {
		return "(None)"
	}
}

// Present returns true if the Optional contains a value.
// Code:
//     some := goptional.New("hello")
//     println(some.Present())
//     none := goptional.None()
//     println(none.Present())
//
// Output:
//     true
//     false
func (o Optional[T]) Present() bool {
	return o.present
}

// Unwrap unwraps the Optional and returns the underlying value, or EmptyOptional as an error if Empty.
// Code:
//     some := goptional.New("hello")
//     println(some.Unwrap())
//     none := goptional.None()
//     println(none.Unwrap())
//
// Output:
//     "hello", nil
//     "", "optional is empty"
func (o Optional[T]) Unwrap() (T, error) {
	if !o.present {
		return o.value, EmptyOptional
	}

	return o.value, nil
}

// UnwrapOr unwraps the Optional and returns the underlying value if Present, or the provided value if Empty.
// Code:
//     some := goptional.New("hello")
//     println(some.UnwrapOr("world"))
//     none := goptional.None()
//     println(none.UnwrapOr("world"))
//
// Output:
//     "hello"
//     "world"
func (o Optional[T]) UnwrapOr(or T) T {
	if !o.present {
		return or
	}

	return o.value
}

// UnwrapOrElse unwraps the Optional and returns the underlying value if Present, or calls the provided function to
// generate a default value if Empty.
//
// Code:
//     makeDefault := func() string {
//         return "world"
//     }
//
//     some := goptional.New("hello")
//     println(some.UnwrapOrElse(makeDefault))
//     none := goptional.None()
//     println(none.UnwrapOrElse(makeDefault))
//
// Output:
//     "hello"
//     "world"
func (o Optional[T]) UnwrapOrElse(orElse func() T) T {
	if !o.present {
		return orElse()
	}

	return o.value
}

// Expect returns the value of the Optional or panics with the provided message if Empty.
// Code:
//     some := goptional.New("hello")
//     println(some.Expect("returns"))
//     none := goptional.None()
//     println(none.Present())
//
// Output:
//     true
//     false
func (o Optional[T]) Expect(msg string) T {
	if !o.present {
		panic(fmt.Errorf("%w: %s", EmptyOptional, msg))
	}
	return o.value
}
