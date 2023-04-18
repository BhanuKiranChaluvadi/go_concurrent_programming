package main

import "fmt"

type order struct {
	ProductCode int
	Quantity    float64
	Status      orderStatus
}

type invalidOrder struct {
	order order
	err   error
}

func (o order) String() string {
	return fmt.Sprintf("ProductCode: %v, Quantity: %v, Status: %v\n", o.ProductCode, o.Quantity, orderStatusToString(o.Status))
}

func orderStatusToString(status orderStatus) string {
	switch status {
	case none:
		return "none"
	case new:
		return "new"
	case received:
		return "received"
	case reserved:
		return "reserved"
	case filled:
		return "filled"
	}
	return "unknown status"
}

type orderStatus int

const (
	none orderStatus = iota
	new
	received
	reserved
	filled
)
