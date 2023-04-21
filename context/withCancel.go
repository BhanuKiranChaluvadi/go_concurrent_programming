package main

/*
import (
	"context"
	"sync"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for range time.Tick(500 * time.Millisecond) {
			if err := ctx.Err(); err != nil {
				println(err.Error())
				return
			}
			println("tick!")
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	wg.Wait()

}

*/
