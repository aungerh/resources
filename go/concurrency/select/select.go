package main

// Note on data structures:
// Firstly, you don't need to pass a pointer to the channel; channels, like maps and others, are references, meaning the underlying data isn't copied, only a pointer to the actual data. If you need a pointer to a chan itself, you'll know when that time comes.

import (
	"fmt"
	"time"
)

// https://golang.org/ref/spec#Select_statements
func main() {
	// A select with a default clause doesn't block
	// A select without a default clause blocks

	// If one or more of the communications can proceed, a single one that can proceed is chosen via a uniform pseudo-random selection.
	// Otherwise, if there is a default case, that case is chosen.
	// If there is no default case, the "select" statement blocks until at least one of the communications can proceed.
	c := make(chan string)

	// this goroutine constantly prints "HI..." to channel "c" every 2 seconds, n times.
	go func(c chan<- string, max int) {
		counter := 0
		for {
			select {
			case _ = <-time.Tick(2 * time.Second):
				if counter == max {
					close(c)
					return
				}
				c <- "HI..."
				counter++
			}
		}
	}(c, 3)

	// if we have a single select with no default outside of an infinite for loop
	// this blocks.
	// select + no_default = block
	// resulting in the program blocked for 2 seconds until the goroutine puts something in the channel
	// prints and exits.

	// select {
	// case message := <-c:
	// 	fmt.Println(message)
	// }

	// if however we have the select with a default clause, the main goroutine will skip the first case and
	// jump directly to the default, as it is the nature of the select statement:
	// "If one or more of the communications can proceed, a single one that can proceed is chosen via a uniform pseudo-random selection"
	// so the following code doesn't block:

	// select {
	// case message := <-c:
	// 	fmt.Println(message)
	// default:
	// 	fmt.Println("too late goroutine...")
	// }

	// this empties the channel
	for v := range c {
		fmt.Println(v)
	}
}
