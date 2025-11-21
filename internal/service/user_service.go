package service

import (
	"fmt"
	"playground/internal/model"
	"playground/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUser() []model.User{
	return s.repo.FindAll()
}

func (s *UserService) CreateUser(name string) model.User{
	user := model.User{
		Name : name,
	}
	return s.repo.Create(user)
}

func(s *UserService) GetUserById(id int) (*model.User, error) {
	user := s.repo.FindByID(id)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *UserService) UpdateUser(id int, name string) (*model.User, error) {
	if name == "" {
		return nil, fmt.Errorf("Name cannot be empty")
	}
	return s.repo.Update(id, name)
}	