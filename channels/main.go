package main

import (
	"fmt"
	"sync"
)

/*
// main function fails to run because of deadlock
func main() {
	ch := make(chan string)
	ch <- "Hello"
	fmt.Println(<-ch)
}
*/

/*
// bidirectional channel
func main() {
	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ch chan string) {
		ch <- "Hello"
	}(ch)
	go func(ch chan string) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
*/

// Directional channel
// It is always the arrow facing towards the left
func main() {
	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)

	// send-only channel
	go func(ch chan<- string) {
		ch <- "Hello"
	}(ch)

	// receive-only channel
	go func(ch <-chan string) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
