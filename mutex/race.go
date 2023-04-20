// go run -race .
package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []int{}

	iterations := 1000

	wg := sync.WaitGroup{}
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			s = append(s, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("length of s: %d\n", len(s))
}
