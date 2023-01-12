package user_utils

import (
	"errors"
	"fmt"
	"golang.com/forum/config"
	"golang.com/forum/models"
	user_domain "golang.com/forum/user/domain"
	"golang.org/x/crypto/bcrypt"
)

func AutoMigrate()  {
	err := config.Connection().AutoMigrate(&user_domain.User{})
	if err != nil {
		return 
	}
}

func Login (email string, password string) (*user_domain.User, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil,err
	}
	err = CheckPassword(user.Password, password)
	if err != nil {
		fmt.Println("Error in password match", err)
		return nil, err
	}
	
	return user, nil
}

func CheckPassword(current string, hash string) error  {
	return bcrypt.CompareHashAndPassword([]byte(current),[]byte(hash))
}

func existEmail(email string) (*bool,error) {
	var us = &user_domain.User{Email: email}
	userGorm := config.DbGorm.First(&us);
	if userGorm.Error != nil{
		fmt.Println(userGorm.Error.Error())
		return nil, userGorm.Error
	}

	exists := userGorm.RowsAffected > 0
	return &exists, nil
}

func SaveUser(user *user_domain.User) error {
	exist, err := existEmail(user.Email)
	if err != nil {
		fmt.Printf("error check user email. %v", err.Error())
		return err
	}
	if *exist == true{
		fmt.Printf("user already exists with email. %v", user.Email)
		return models.ErrorUserRegistred
	}

	result := config.DbGorm.Create(&user);
	
	if result.Error != nil {
		fmt.Printf("error saving user. %v", result.Error.Error())
	}
	fmt.Printf("saved user %v.", user.Id)

	return nil
}



func GetUserByEmail(email string) (*user_domain.User, error) {
	var us = &user_domain.User{Email: email}
	userGorm := config.DbGorm.First(&us, "email = ?", email);
	if userGorm.Error != nil{
		return nil, userGorm.Error
	}

	if userGorm.RowsAffected > 0 {
		fmt.Printf("user rows %v", us)
		fmt.Printf("found rows %v", userGorm.RowsAffected)
		return us, nil
	}

	fmt.Printf("not found rows %v", userGorm.RowsAffected)

	return nil, errors.New("not found rows.")
}


