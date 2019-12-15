# Go Lessons

- [Complex types (such as array, slices, maps, channels, struct...) are **references**. And therefore when calling a function with any of these as an argument, any modifications to it from the body of it will change the original value](https://github.com/refs/resources/blob/master/go/types/complex_types_references/main.go)
  - [If a map isn't a reference variable, what is it?](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it)
  - [Are maps passed by value or by reference in Go?](https://stackoverflow.com/questions/40680981/are-maps-passed-by-value-or-by-reference-in-go)
  - Mind then that, as the code implies, when passing a map / channel to a function, this is a reference.
