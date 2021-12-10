package storage

import (
	"fmt"
	"learngo/api_only_go_task/pkg/model"
	"log"
)

var users = make(map[string]*model.User)

type UserStorage interface {
	SaveUser(user *model.User) error
	GetUser(email string) (*model.User, error)
}
type userStorage struct{}

func NewUserStorage() *userStorage {
	return &userStorage{}
}

func (s *userStorage) SaveUser(user *model.User) error {
	if _, ok := users[user.Email]; !ok {
		users[user.Email] = user
	} else {
		log.Println(fmt.Sprintf("User %s already stored", user.Email))
		return fmt.Errorf("User %s already stored", user.Email)
	}

	return nil
}

func (s *userStorage) GetUser(email string) (*model.User, error) {
	if userData, ok := users[email]; ok {
		return userData, nil
	} else {
		return nil, fmt.Errorf("The User doesn't  exist ")
	}

}
