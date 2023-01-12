package user_application

import (
	"fmt"
	user_domain "golang.com/forum/user/domain"
	user_utils "golang.com/forum/user/utils"
)


func CreateUser(user *user_domain.User) (*user_domain.User, error) {
	userCreated, err := user_domain.NewUser(user.Name, user.LastName, user.Password, user.Email)
	if err != nil {
		return nil, err
	}
	
	saveUser := user_utils.SaveUser(userCreated)
	if saveUser != nil {
		fmt.Println("Unexpected error trying to save user. ", saveUser.Error())
		return nil, saveUser
	}
	return user, nil 

}
