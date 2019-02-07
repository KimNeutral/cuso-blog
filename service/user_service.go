package service

import (
	"cuso-blog/model"
	"cuso-blog/repository"
)

type UserSerivce struct {
	Repository *repository.UserRepository
}

// GetUserByEmail : Retrive User data with email
func (s *UserSerivce) GetUserByEmail(email string) (user model.User, err error) {
	user, err = s.Repository.GetUserByEmail(email)
	return
}

// GetUserByID : Retrive User data with id
func (s *UserSerivce) GetUserByID(id model.Primary) (user model.User, err error) {
	user, err = s.Repository.GetUserByID(id)
	return
}

func NewUserService(repository *repository.UserRepository) *UserSerivce {
	instance := new(UserSerivce)

	instance.Repository = repository

	return instance
}
