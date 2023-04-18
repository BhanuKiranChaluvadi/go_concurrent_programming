package main

func main() {

	ch := make(chan string, 3)

	for _, v := range []string{"a", "b", "c"} {
		ch <- v
	}
	// closes inlet of the channel
	close(ch)

	for msg := range ch {
		println(msg)
	}

}
