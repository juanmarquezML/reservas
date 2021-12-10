package storage

import (
	"fmt"
	"learngo/api_only_go_task/pkg/model"
	"log"
)

var customers = make(map[string]*model.Customer)

type CustomerStorage interface {
	SaveCustomer(cts *model.Customer) error
	//GetCustomer(email string) (*model.User, error)
}
type customerStorage struct{}

func NewCustomerStorage() *customerStorage {
	return &customerStorage{}
}

func (s *customerStorage) SaveCustomer(cts *model.Customer) error {
	if _, ok := customers[cts.Cuit]; !ok {
		customers[cts.Cuit] = cts
	} else {
		log.Println(fmt.Sprintf("Customer %s already stored", cts.Name))
		return fmt.Errorf("Customer %s already stored", cts.Name)
	}

	return nil
}

/*
func (s *userStorage) GetUser(email string) (*model.User, error) {
	if userData, ok := users[email]; ok {
		return userData, nil
	} else {
		return nil, fmt.Errorf("The User doesn't  exist ")
	}

}*/
