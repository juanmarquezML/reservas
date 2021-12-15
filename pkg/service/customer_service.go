package service

import (
	"learngo/api_only_go_task/pkg/model"
	"learngo/api_only_go_task/pkg/storage"
)

type CustomerService interface {
	CreateCustomer(custumer *model.Customer) error
	GetCustomer(cuit string) (*model.Customer, error)
}

type customerService struct {
	customerStorage storage.CustomerStorage
}

func NewCustomerService(u storage.CustomerStorage) *customerService {
	return &customerService{u}
}

func (s *customerService) CreateCustomer(cts *model.Customer) error {
	return s.customerStorage.SaveCustomer(cts)
}

func (s *customerService) GetCustomer(cuit string) (*model.Customer, error) {
	cts, error := s.customerStorage.GetCustomer(cuit)
	return cts, error
}
