package app

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email          string `json:"email,omitempty"`
	Password       string `json:"password,omitempty"`
	ActivationLink string `json:"activation_link,omitempty"`
	ID             int    `json:"id,omitempty"`
	IsActivated    bool   `json:"is_activated,omitempty"`
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 3)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

func (u *User) GenerateActivationLink() {
	u.ActivationLink = uuid.NewString()
}
