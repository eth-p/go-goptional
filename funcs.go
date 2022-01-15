package goptional

// Map maps an Optional[T] to an Optional[U] by applying a function to the value (if Present).
// If no value is present, the Optional[U] will be Empty.
//
// Code:
//     some := goptional.New[int](5)
//     println(some.Expect("has value"))
//     someAsString := goptional.Map(some, func(value int) string {
//         return fmt.Sprintf("the number is %v", value)
//     })
//     println(some.Expect("has value"))
//
// Output:
//     5
//     "the number is 5"
func Map[T any, U any](optional Optional[T], mapper func(from T) U) Optional[U] {
	return MapInto(optional, func(from T) Optional[U] {
		return New[U](mapper(from))
	})
}

// MapInto maps an Optional[T] to an Optional[U] by applying a function to the value (if Present).
//
// This differs from the regular Map function by allowing the mapping function to return an Optional[U].
// If an Optional[U] is provided by the mapping function, an Empty Optional[U] will be returned.
//
// Code:
//     some := goptional.New[int](5)
//     println(some.Expect("has value"))
//     someAsString := goptional.MapInto(some, func(value int) goptional.Optional[string] {
//         return goptional.None()
//     })
//     println(some.Expect("will panic"))
//
// Output:
//     5
//     <panic>
func MapInto[T any, U any](optional Optional[T], mapper func(from T) Optional[U]) Optional[U] {
	value, err := optional.Unwrap()
	if err != nil {
		return None[U]()
	}

	return mapper(value)
}
