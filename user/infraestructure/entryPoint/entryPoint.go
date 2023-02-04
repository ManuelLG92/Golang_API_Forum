package entryPoint

import (
	"errors"
	"fmt"
	"forum/config"
	userDomain "forum/user/domain"
)

type UserDto struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func FetchUserById(id string) (*UserDto, error) {
	var user = &userDomain.User{}
	userGorm := config.DbGorm.First(&user, "id = ?", id)
	if userGorm.Error != nil || user.Id == "" {
		fmt.Println(userGorm.Error.Error())
		fmt.Println(user.Id)
		return nil, errors.New(fmt.Sprintf("User %v not found", id))
	}
	return &UserDto{Id: user.Id, Name: user.Name, LastName: user.LastName}, nil
}
