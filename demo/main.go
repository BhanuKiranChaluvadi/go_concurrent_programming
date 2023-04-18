package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	receiveOrders()
	fmt.Println(orders)
}

func receiveOrders() {
	for _, rawOrder := range rawOrder {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Printf("Error unmarshalling order: %v", err)
			continue
		}
		orders = append(orders, newOrder)
	}
}

var rawOrder = []string{
	`{"ProductCode": 1111, "Quantity": 5, "Status": 1}`,
	`{"ProductCode": 2222, "Quantity": 42.3, "Status": 1}`,
	`{"ProductCode": 3333, "Quantity": 19, "Status": 1}`,
	`{"ProductCode": 4444, "Quantity": 8, "Status": 1}`,
}
