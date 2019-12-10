package main

import (
	"fmt"
	"time"
)

/*
	When a channel is closed and there's a goroutine listening on it
	it keeps yielding zero values of the channel's type.
*/

func main() {
	c := make(chan int)
	go printEverySecond(c)

	// For channels, the iteration values produced are the successive values sent on the channel until the channel is closed. If the channel is nil, the range expression blocks forever.
	// src: https://golang.org/ref/spec
	for r := range c {
		fmt.Println(r)
	}

	// Even when the channel is closed it keeps emiting zero values (0)
	// for {
	// 	select {
	// 	case v := <-c:
	// 		fmt.Println(v)
	// 	}
	// }
}

func printEverySecond(c chan<- int) {
	for i := 0; i < 4; i++ {
		time.Sleep(1 * time.Second)
		c <- i
	}

	// I'm done writing to the channel so I close it
	close(c)
}
