package service

import (
	"learngo/api_only_go_task/pkg/model"
	"learngo/api_only_go_task/pkg/storage"
)

type CustumerService interface {
	CreateCustumer(custumer *model.Customer) error
	//GetCustumer(cuit string) (*model.Customer, error)
}

type customerService struct {
	customerStorage storage.CustomerStorage
}

func NewCustomerService(u storage.CustomerStorage) *customerService {
	return &customerService{u}
}

func (s *customerService) CreateCustumer(cts *model.Customer) error {
	return s.customerStorage.SaveCustomer(cts)
}

/*func (s *userService)GetUser(email string) (*model.User, error) {
	user, error:=s.userStorage.GetUser(email)
	return user, error
}*/
