package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan string, 1)

	ch1 <- 999
	// ch2 <- "Hello" // behavious is undefined when both channels are ready

	// unlike switch, select will not check in order of the cases.
	// It will check if any of the cases are ready to be executed.
	// if both of the cases are ready, it will execute the randome case.
	// the runtime actually will randomly select one of the cases to execute.
	// some time we get the message 999 and some time we get the message "Hello"

	// if both the channels are commented out, the select will block forever
	// because there is no case that is ready to be executed. Its a deadlock.
	// so add default case to avoid deadlock

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("No message received")
	}
}
