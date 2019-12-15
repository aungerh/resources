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
	printMem(a)
	fmt.Print(a)
}

func printMem(a map[string]int) {
	fmt.Printf("%p\n", a)
	a["a"] = 3 // this modifies the original variable. maps are reference types
	// Arguably renaming the type from *map[int]int to map[int]int, while confusing because the
	// type does not look like a pointer, was less confusing than a pointer shaped value which
	// cannot be dereferenced. (2)
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
