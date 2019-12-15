# Go Lessons

- Complex types (array, slice, map, channel) are **references** [1]
- When passing a complex type to a function, changes in the body of this will alter its original value
- Zero values of map, channel, slice, interfaces, pointers and functions is <nil>[2]

- [1](https://github.com/refs/resources/blob/master/go/types/complex_types_references/main.go)
  - [If a map isn't a reference variable, what is it?](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it)
  - [Are maps passed by value or by reference in Go?](https://stackoverflow.com/questions/40680981/are-maps-passed-by-value-or-by-reference-in-go)
  - Mind then that, as the code implies, when passing a map / channel to a function, this is a reference.
- [2](https://github.com/refs/resources/blob/master/go/types/complex_types_references/main.go#L72)