package customers

import (
	"fmt"
	"modules-go/orders"
)

type Customer struct {
	Name          string
	Surname       string
	Age           int
	FavoriteOrder orders.Order
}

func (customer Customer) FullName() string {
	return fmt.Sprintf("%s %s", customer.Name, customer.Surname)
}
