package mapstore

import (
	"customerapp/domain"
	"fmt"

	"github.com/gofrs/uuid"
)

type MapStore struct {
	store map[string]domain.Customer
}

func NewMapStore() (*MapStore, error) {
	return &MapStore{store: make(map[string]domain.Customer)}, nil
}

func (mapstore *MapStore) Create(cust domain.Customer) error {
	if _, ok := mapstore.store[cust.CustomerID]; ok {
		return domain.ErrCustomerExists
	}

	uid, _ := uuid.NewV4()
	cust.ID = uid.String()
	mapstore.store[cust.CustomerID] = cust
	return nil
}

func (mapstore *MapStore) Update(id string, customer domain.Customer) error {

	if _, ok := mapstore.store[id]; !ok {
		return domain.ErrNotFound
	}

	mapstore.store[id] = customer
	return nil
}

// GetById(string) (Customer, error)
func (mp *MapStore) GetById(id string) (domain.Customer, error) {
	fmt.Println("Called with id:%s", id)

	if cust, ok := mp.store[id]; !ok {
		return domain.Customer{}, domain.ErrNotFound
	} else {
		return cust, nil
	}
}

// func (i *MapStore) GetById(id string) (domain.Customer, error) {
// 	if v, ok := i.store[id]; !ok {
// 		return domain.Customer{}, domain.ErrNotFound
// 	} else {
// 		return v, nil
// 	}

// }
func (mp *MapStore) Delete(id string) error {
	if _, ok := mp.store[id]; !ok {

		return domain.ErrNotFound
	}

	delete(mp.store, id)
	return nil
}
func (mp *MapStore) GetAll() ([]domain.Customer, error) {

	customers := make([]domain.Customer, 0)
	for _, customer := range mp.store {
		customers = append(customers, customer)
	}
	return customers, nil
}
