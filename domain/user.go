package domain

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Base
	Name string `json:"name" gorm:"type:vachar(255)"`
	Email string `json:"name" gorm:"type:vachar(255);unique_index"`
	Password string  `json:"-" gorm:"type:vachar(255)"`
	Token string  `json:"token" gorm:"type:vachar(255);unique_index"`
}

func (user *User) Prepare() error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Error during the password generation: %v", err)
		return err
	}

	user.Password = string(password)
	user.Token = uuid.NewV4().String()

	err = user.validate()
	
	if err != nil {
		log.Fatalf("Error during the user validation: %v", err)
		return err
	}

	return nil

}

func (user *User) validate() error {
	return nil
}