package user_domain

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.com/forum/handlers"
	"golang.com/forum/helpers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

)


type User struct {
	gorm.Model
	Id string `json:"id" sql:"type:VARCHAR(255)" json:"id" gorm:"primaryKey"`
	Name string `json:"name" validate:"required" sql:"type:VARCHAR(100)"`
	LastName string `json:"last_name" sql:"type:VARCHAR(100)" gorm:"column:last_name" validate:"required"`
	Email string `json:"email" sql:"type:VARCHAR(150);not null;unique" validate:"required,email"`
	Password string `json:"password" sql:"type:VARCHAR(255)" validate:"required,gte=8,lte=64"`
}

func (user *User) IsValidOnCreation() bool {
	if user.Email == "" || user.Name == "" || user.LastName == "" || user.Password == "" {
		return false
	}
	return true
}

func (user *User) Valid() error {

	if handlers.ValidEmail(user.Email) != nil {
		return helpers.ErrorEmail
	}
	if user.validData() != nil {
		return helpers.ErrorNotValidData
	}
	return nil
}
func (user *User)validData() error  {
	if len(user.Name) < 1 ||  len(user.LastName) > 30 {
		fmt.Println("Name invalid")
		return helpers.ErrorEmptyUsername

	}
	if len(user.LastName) < 1 || len(user.LastName) > 30 {
		fmt.Println("Lastname invalid")
		return helpers.ErrorLastname

	}
	if len(user.Password) < 8 || len(user.Password) > 24 {
		fmt.Println("Password invalid")
		return helpers.ErrorPassword

	}
	return nil
}

func NewUser(name, lastName, password, email string) (*User, error) {
	user := &User{Id: uuid.New().String(), Name: name, LastName: lastName, Password: password, Email: email} // Creamos un objeto user o instancia en Java referenciando con &
	if err := setPassword(user); err != nil {
		return nil,err
	}
	if user.IsValidOnCreation() == false {
		return nil, errors.New("USer is not valid")
	}
	return user, nil
}

func setPassword (user *User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return helpers.ErrorPasswordEncryption
	}
	user.Password = string(hash)
	return nil
}