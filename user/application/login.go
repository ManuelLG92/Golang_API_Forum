package user_application

import (
	"errors"
	"fmt"
	"golang.com/forum/auth"
	user_utils "golang.com/forum/user/utils"
)

func Login(email, password, ip string) (*string, error)  {
	user, err := user_utils.Login(email, password)
	if err != nil {
		fmt.Println("Password and/or email invalid", err.Error())
		return nil, errors.New(fmt.Sprintf("Password and/or email invalid. %v", err.Error()))
	}

	err, token := auth.GenerateJwt(auth.JwtCustomClaims{Id: user.Id, Name: user.Name, Email: user.Email, Password: user.Password, IP: ip})
	if err != nil {
		fmt.Printf("Error creating token. %v", err.Error())
		return nil, errors.New(fmt.Sprintf("Unexpected error. %v", err.Error()))
	}
	return token, err
}