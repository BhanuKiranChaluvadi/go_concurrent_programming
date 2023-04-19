package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// var receivedOrdersCh = make(chan order)
	// var validatedOrdersCh = make(chan order)
	// var invalidOrdersCh = make(chan invalidOrder)

	// go receiveOrders(receivedOrdersCh)
	receivedOrdersCh := receiveOrders()
	validatedOrdersCh, invalidOrdersCh := validateOrders(receivedOrdersCh)

	wg.Add(1)
	go func(validatedOrdersCh <-chan order, invalidOrdersCh <-chan invalidOrder) {
	loop:
		for {
			select {
			case order, ok := <-validatedOrdersCh:
				if ok {
					fmt.Printf("Valid order received: %v", order)
				} else {
					break loop
				}

			case invalidOrder, ok := <-invalidOrdersCh:
				if ok {
					fmt.Printf("Invalid order received: %v. Issue: %v\n", invalidOrder.order, invalidOrder.err)
				} else {
					break loop
				}
			}
		}
		wg.Done()
	}(validatedOrdersCh, invalidOrdersCh)

	wg.Wait()
}

func validateOrders(in <-chan order) (<-chan order, <-chan invalidOrder) {
	out := make(chan order)
	errCh := make(chan invalidOrder, 1)
	go func(out chan<- order, errCh chan<- invalidOrder) {
		for order := range in {
			if order.Quantity < 0 {
				errCh <- invalidOrder{order: order, err: fmt.Errorf("invalid quantity: %v. Quantity must be greater than zero", order.Quantity)}
			} else {
				out <- order
			}
		}
		close(out)
		close(errCh)
	}(out, errCh)
	return out, errCh
}

func receiveOrders() <-chan order {
	out := make(chan order)
	go func(out chan<- order) {
		for _, rawOrder := range rawOrder {
			var newOrder order
			err := json.Unmarshal([]byte(rawOrder), &newOrder)
			if err != nil {
				log.Printf("Error unmarshalling order: %v", err)
				continue
			}
			out <- newOrder
		}
		close(out)
	}(out)
	return out
}

var rawOrder = []string{
	`{"ProductCode": 1111, "Quantity": 5, "Status": 1}`,
	`{"ProductCode": 2222, "Quantity": 42.3, "Status": 1}`,
	`{"ProductCode": 3333, "Quantity": 19, "Status": 1}`,
	`{"ProductCode": 4444, "Quantity": 8, "Status": 1}`,
}
