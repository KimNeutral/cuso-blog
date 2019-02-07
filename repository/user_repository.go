package repository

import (
	"cuso-blog/model"
	"log"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	Repository
}

// GetUserByEmail : Retrive User data with email
func (r *UserRepository) GetUserByEmail(email string) (user model.User, err error) {
	r.DB.Where(model.User{Email: email}).First(&user)
	return
}

// GetUserByID : Retrive User data with id
func (r *UserRepository) GetUserByID(id model.Primary) (user model.User, err error) {
	r.DB.Where(model.User{Model: gorm.Model{ID: uint(id)}}).Find(&user)

	log.Printf("Name: " + user.Name)
	return
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	instance := new(UserRepository)

	instance.DB = db

	return instance
}
