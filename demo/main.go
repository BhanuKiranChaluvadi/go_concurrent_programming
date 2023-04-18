package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go receiveOrders(&wg)
	wg.Wait()
	fmt.Println(orders)
}

func receiveOrders(wg *sync.WaitGroup) {
	for _, rawOrder := range rawOrder {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Printf("Error unmarshalling order: %v", err)
			continue
		}
		orders = append(orders, newOrder)
	}
	wg.Done()
}

var rawOrder = []string{
	`{"ProductCode": 1111, "Quantity": 5, "Status": 1}`,
	`{"ProductCode": 2222, "Quantity": 42.3, "Status": 1}`,
	`{"ProductCode": 3333, "Quantity": 19, "Status": 1}`,
	`{"ProductCode": 4444, "Quantity": 8, "Status": 1}`,
}
