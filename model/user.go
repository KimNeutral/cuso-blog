package model

import "github.com/jinzhu/gorm"

// User : user table
type User struct {
	gorm.Model
	Name  string
	Email string
}

// UserInfo : table holds user's info
type UserInfo struct {
	GithubURL   string
	FacebookURL string
	TwitterURL  string
}

// Users : plural type of User
type Users []User

// UsersInfo : plural type of UserInfo
type UsersInfo []UserInfo
