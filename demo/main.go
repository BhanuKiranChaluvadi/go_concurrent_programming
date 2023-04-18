package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var receivedOrdersCh = make(chan order)
	var validatedOrdersCh = make(chan order)
	var invalidOrdersCh = make(chan invalidOrder)

	go receiveOrders(receivedOrdersCh)
	go validateOrders(receivedOrdersCh, validatedOrdersCh, invalidOrdersCh)

	wg.Add(1)
	go func() {
		order := <-validatedOrdersCh
		fmt.Printf("Valid order received: %v", order)
		wg.Done()
	}()

	go func() {
		invalidOrder := <-invalidOrdersCh
		fmt.Printf("Invalid order received: %v. Issue: %v\n", invalidOrder.order, invalidOrder.err)
		wg.Done()
	}()
	wg.Wait()
}

func validateOrders(in, out chan order, errCh chan invalidOrder) {
	order := <-in
	if order.Quantity < 0 {
		errCh <- invalidOrder{order: order, err: fmt.Errorf("invalid quantity: %v. Quantity must be greater than zero", order.Quantity)}
	} else {
		out <- order
	}
}

func receiveOrders(out chan order) {
	for _, rawOrder := range rawOrder {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Printf("Error unmarshalling order: %v", err)
			continue
		}
		out <- newOrder
	}
}

var rawOrder = []string{
	`{"ProductCode": 1111, "Quantity": 5, "Status": 1}`,
	`{"ProductCode": 2222, "Quantity": 42.3, "Status": 1}`,
	`{"ProductCode": 3333, "Quantity": 19, "Status": 1}`,
	`{"ProductCode": 4444, "Quantity": 8, "Status": 1}`,
}
