package user_utils

import (
	"errors"
	"fmt"
	"forum/config"
	user_domain "forum/user/domain"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (*user_domain.User, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	err = CheckPassword(user.Password, password)
	if err != nil {
		fmt.Println("Error in password match", err)
		return nil, err
	}

	return user, nil
}

func CheckPassword(current string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(current), []byte(hash))
}

func existEmail(email string) bool {
	var us = &user_domain.User{}
	userGorm := config.DbGorm.First(&us, "email = ?", email)
	if userGorm.Error != nil {
		fmt.Printf("erro exist email. %v", userGorm.Error.Error())
		fmt.Println()
		return false
	}
	return true
}

func SaveUser(user *user_domain.User) error {
	var userExists = existEmail(user.Email)
	fmt.Printf("exists?. %v", userExists)
	if userExists == true {
		fmt.Printf("user already exists with email. %v", user.Email)
		return errors.New(fmt.Sprintf("user already exists with email. %v", user.Email))
	}

	result := config.DbGorm.Create(&user)

	if result.Error != nil {
		fmt.Printf("error saving user. %v", result.Error.Error())
	}
	fmt.Printf("saved user %v.", user.Id)

	return nil
}

func GetUserByEmail(email string) (*user_domain.User, error) {
	var user = &user_domain.User{}
	dbResult := config.DbGorm.First(&user, "email = ?", email)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	if dbResult.RowsAffected > 0 {
		fmt.Printf("user rows %v", user)
		fmt.Printf("found rows %v", dbResult.RowsAffected)
		return user, nil
	}

	fmt.Printf("not found rows %v", dbResult.RowsAffected)

	return nil, errors.New("not found rows")
}
