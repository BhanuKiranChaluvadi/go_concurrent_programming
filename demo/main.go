package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	receivedOrdersCh := receiveOrders()
	validatedOrdersCh, invalidOrdersCh := validateOrders(receivedOrdersCh)
	reservedInventoryCh := reserveInventory(validatedOrdersCh)
	fillOrders(reservedInventoryCh, &wg)

	wg.Add(1)

	go func(invalidOrdersCh <-chan invalidOrder) {
		for order := range invalidOrdersCh {
			fmt.Printf("Invalid order: %v, Error: %v\n", order.order, order.err)
		}
		wg.Done()
	}(invalidOrdersCh)

	wg.Wait()
}

func fillOrders(in <-chan order, wg *sync.WaitGroup) {

	const workers = 3
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(in <-chan order) {
			for order := range in {
				order.Status = filled
				fmt.Printf("Order has been completed: %v\n", order)
			}
			wg.Done()
		}(in)
	}
}

func reserveInventory(in <-chan order) <-chan order {
	out := make(chan order)

	var wg sync.WaitGroup
	const workers = 3

	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(in <-chan order, out chan<- order) {
			for order := range in {
				order.Status = reserved
				out <- order
			}
			wg.Done()
		}(in, out)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
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
	`{"ProductCode": 3333, "Quantity": -19, "Status": 1}`,
	`{"ProductCode": 4444, "Quantity": 8, "Status": 1}`,
}
