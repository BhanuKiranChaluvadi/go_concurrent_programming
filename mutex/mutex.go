package main

/*
// Problem: The length of s is not 1000
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
	fmt.Printf("length of s: %d", len(s))
}
*/

/*
func main() {
	s := []int{}
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	const iterations = 1000
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			s = append(s, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("length of s: %d\n", len(s))
}
*/
