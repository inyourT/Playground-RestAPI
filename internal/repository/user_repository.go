package repository

import (
	"fmt"
	"playground/internal/model"
)

type UserRepository struct {
	users []model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users : []model.User{
			{ID:1, Name:"Rusman"},
			{ID:2, Name:"Ramzi"},
		},
	}
}

func (r *UserRepository) FindAll() []model.User{
	return r.users
}

func (r *UserRepository) Create(user model.User) model.User{
	// Auto-increment
	newId := len(r.users)+1
	user.ID = newId

	r.users = append(r.users, user)

	return user
}

func (r *UserRepository) FindByID(id int) *model.User{
	for _, u := range r.users{
		if u.ID == id {
			return &u
		}
	}
	return nil
}

func (r *UserRepository) Update(id int, name string) (*model.User, error){
	for i, u := range r.users{
		if u.ID == id {
			r.users[i].Name = name
			return &r.users[i], nil
		}
	}
	return nil, fmt.Errorf("User not found")
}
