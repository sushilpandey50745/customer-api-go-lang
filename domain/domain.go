package domain

import "errors"

var ErrNotFound = errors.New("No Customer record Found")
var ErrCustomerExists = errors.New("Customer with given CustID exists")

type Customer struct {
	ID         string `json:"id, omitempty"`
	CustomerID string `json:"custid"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

type Customerstore interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetById(string) (Customer, error)
	GetAll() ([]Customer, error)
}
