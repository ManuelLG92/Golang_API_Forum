package handlers

import (
	"fmt"
	"golang.com/forum/config"
	"golang.com/forum/models"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

var emailRegexp = regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])")

//Init Login
func login (email string, password string) (*User, error) {
	user, _ := GetUserByEmail(email)
	//fmt.Println(user)
	//fmt.Println(password)
	err := CheckPassword(user.Password, password)
	if err != nil {
		fmt.Println("Error in password match", err)
		return nil, err
	}
	
	return user, nil

}
// End Login


// Init ValidEmail
func validEmail (email string) error {
	if !emailRegexp.MatchString(email) {
		return models.ErrorEmail
	}
	return nil

}
// End ValidEmail

// Init validData
func validData(name, lastName, password string) error  {
	if len(name) < 1 ||  len(name) > 30 {
		fmt.Println("NAme invalid")
		return models.ErrorEmptyUsername

	}
	if len(lastName) < 1 || len(lastName) > 30 {
		fmt.Println("Lastname invalid")
		return models.ErrorLastname

	}
	if len(password) < 8 || len(password) > 24 {
		fmt.Println("Password invalid")
		return models.ErrorPassword

	}
	return nil
}
// End validData

//Init Encrypt user password
func (user *User) SetPassword (password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return models.ErrorPasswordEncryption
	}
	user.Password = string(hash)
	return nil
}
// End Encrypt user password

// Init Post validations

// Init user id
func validUserId ( id string) error {
	sql := "SELECT * FROM users where id=?"
	_ , err := config.Query(sql, id)
	if err != nil {
		return models.ErrorPostByUserId
	}
	return nil

}
// End user id

func validPostData (title, content string) error{
	if len(title) < 1 || len(title) > 50 || len(content) < 1 || len(content) > 255{
		return models.ErrorPostData
	}
	return nil
}

//End Post validations