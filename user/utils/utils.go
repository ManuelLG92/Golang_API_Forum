package user_utils

import (
	"errors"
	"fmt"
	"forum/config"
	"forum/storage"
	userDomain "forum/user/domain"
	"golang.org/x/crypto/bcrypt"
)

var cache = storage.NewCache[userDomain.User]()

func Login(email string, password string) (*userDomain.User, error) {
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
	var us = &userDomain.User{}
	if result := cache.Get(email); result != nil {
		return true
	}
	userGorm := config.DbGorm.First(&us, "email = ?", email)
	if userGorm.Error != nil || us.Id == "" {
		fmt.Printf("erro exist email. %v", userGorm.Error.Error())
		fmt.Println()
		return false
	}
	cache.Set(us.Id, *us)
	return true
}

func SaveUser(user *userDomain.User) error {
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

func GetUserByEmail(email string) (*userDomain.User, error) {
	var user = &userDomain.User{}
	if result := cache.Get(email); result != nil {
		return result, nil
	}
	dbResult := config.DbGorm.First(&user, "email = ?", email)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	if dbResult.RowsAffected > 0 {
		fmt.Printf("user rows %v", user)
		fmt.Printf("found rows %v", dbResult.RowsAffected)
		cache.Set(email, *user)
		return user, nil
	}

	fmt.Printf("not found rows %v", dbResult.RowsAffected)

	return nil, errors.New("not found rows")
}
