package customers

import "fmt"

type Customer struct {
	Name    string
	Surname string
	Age     int
}

func (customer Customer) FullName() string {
	return fmt.Sprintf("%s %s", customer.Name, customer.Surname)
}
