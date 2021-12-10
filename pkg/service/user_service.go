package service

import (
	"learngo/api_only_go_task/pkg/model"
	"learngo/api_only_go_task/pkg/storage"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetUser(email string) (*model.User, error)
}
type userService struct {
	userStorage storage.UserStorage
}

func NewUserService(u storage.UserStorage) *userService {
	return &userService{u}
}

func (s *userService)CreateUser(user *model.User) error {
	return s.userStorage.SaveUser(user)
}
func (s *userService)GetUser(email string) (*model.User, error) {
	user, error:=s.userStorage.GetUser(email)
	return user, error
}