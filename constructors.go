package goptional

// From creates a new Optional from a (value, bool) tuple.
// If the bool is false, a None optional will be returned.
//
// You can check if the Optional contains a value using the Present function.
//
// Code:
//     some := goptional.From("hello", true)
//     println(some.Present())
//     println(some.Expect("has a value"))
//     none := goptional.From("goodbye", false)
//     println(none.Present())
//     println(none.Expect("this will panic"))
//
// Output:
//     true
//     hello
//     false
//     <panic>
func From[T any](value T, ok bool) Optional[T] {
	if !ok {
		return None[T]()
	}
	return New[T](value)
}

// FromError creates a new Optional from a (value, error) tuple.
// If the error is not nil, a None optional will be returned.
//
// You can check if the Optional contains a value using the Present function.
//
// Code:
//     some := goptional.FromError("hello", nil)
//     println(some.Present())
//     println(some.Expect("has a value"))
//     none := goptional.FromError("goodbye", errors.New("an error"))
//     println(none.Present())
//     println(none.Expect("this will panic"))
//
// Output:
//     true
//     hello
//     false
//     <panic>
func FromError[T any](value T, err error) Optional[T] {
	if err != nil {
		return None[T]()
	}
	return New[T](value)
}

// FromPtr creates a new Optional from a possibly-nil pointer.
// If the pointer is nil, a None optional will be returned.
//
// You can check if the Optional contains a value using the Present function.
//
// Code:
//     str := "hello"
//     var strPtr *string = &str
//     var nilPtr *string = nil
//     some := goptional.FromPtr(strPtr)
//     println(some.Present())
//     println(some.Expect("has a value"))
//     none := goptional.FromPtr(nilPtr)
//     println(none.Present())
//     println(none.Expect("this will panic"))
//
// Output:
//     true
//     hello
//     false
//     <panic>
func FromPtr[T any](value *T) Optional[*T] {
	if value == nil {
		return None[*T]()
	}
	return New[*T](value)
}

// New creates a new Optional from a value.
//
// Code:
//     some := goptional.New("hello")
//     println(some.Present())
//     println(some.Expect("has a value"))
//
// Output:
//     true
//     hello
func New[T any](value T) Optional[T] {
	return Optional[T]{
		present: true,
		value:   value,
	}
}

// None creates an empty Optional.
//
// Code:
//     some := goptional.New("hello")
//     println(some.Present())
//     println(some.Expect("this will panic"))
//
// Output:
//     false
//     <panic>
func None[T any]() Optional[T] {
	return Optional[T]{}
}
