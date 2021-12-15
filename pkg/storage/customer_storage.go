package storage

import (
	"fmt"
	"learngo/api_only_go_task/pkg/model"
)

var customers = make(map[string]*model.Customer)

type CustomerStorage interface {
	SaveCustomer(cts *model.Customer) error
	GetCustomer(cuit string) (*model.Customer, error)
}
type customerStorage struct{}

func NewCustomerStorage() *customerStorage {
	return &customerStorage{}
}

func (s *customerStorage) SaveCustomer(cts *model.Customer) error {
	if _, ok := customers[cts.Cuit]; !ok {
		customers[cts.Cuit] = cts
	} else {
		return fmt.Errorf("Customer %s already stored", cts.Name)
	}

	return nil
}

func (s *customerStorage) GetCustomer(cuit string) (*model.Customer, error) {
	if ctsData, ok := customers[cuit]; ok {
		return ctsData, nil
	} else {
		return nil, fmt.Errorf("The Customer doesn't  exist ")
	}

}
