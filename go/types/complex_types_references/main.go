package main

import "fmt"

// (1) https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go
// (2) https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it

// "It is not possible to create a Go program where two variables share the same storage location in memory.
// It is possible to create two variables whose contents point to the same storage location,
// but that is not the same thing as two variables who share the same storage location." (1)

func main() {
	a := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Printf("%p\n", a)
	pmemMap(a) // modifies a (passes by reference)

	b := []int{1, 2, 3, 4}
	fmt.Printf("%p\n", b)
	pmemSlice(b) // modifies b (passes by reference)

	c := 42
	fmt.Printf("%p\n", &c)
	pmemInt(c) // doesn't modify c (passes by value)
}

func pmemMap(a map[string]int) {
	fmt.Printf("%p\n", a)
	a["a"] = 3 // this modifies the original variable. maps are reference types
	// Arguably renaming the type from *map[int]int to map[int]int, while confusing because the
	// type does not look like a pointer, was less confusing than a pointer shaped value which
	// cannot be dereferenced. (2)
}

func pmemSlice(b []int) {
	fmt.Printf("%p\n", b)
	b[0] = 3
}

func pmemInt(c int) {
	fmt.Printf("%p\n", &c)
	c = 3
}

// EXTRACT FROM THE BLOG POST (1)
// package main

// import "fmt"

// func main() {
//         var a int
//         var b, c = &a, &a
//         fmt.Println(b, c)   // 0x1040a124 0x1040a124
//         fmt.Println(&b, &c) // 0x1040c108 0x1040c110
// }

// In this example, b and c hold the same value–the address of a–however,
// b and c themselves are stored in unique locations. Updating the contents of b would have no effect on c.
