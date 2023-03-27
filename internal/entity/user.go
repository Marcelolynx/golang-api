package entity

import (
	"github.com/marcelolynx/golang-api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"-"`
	Email    string    `json:"email"`
}

func NewUser(username, password, email string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Username: username,
		Password: string(hash),
		Email:    email,
	}, nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
