package handlers

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.com/forum/config"
)

type SessionsData struct {
	Uuid string `json:"uuid"`
	UserId int `json:"user_id"`
}

func AutoMigrate()  {
	err := config.Connection().AutoMigrate(&User{})
	if err != nil {
		return 
	}
}

func GetUserByEmail(email string) (*User, error) {
	var us = &User{Email: email}
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

	return nil, errors.Errorf("not found rows.")
}

func (user *User) insertUser() error {
	result := config.DbGorm.Create(&user);

	if result.Error != nil {
		fmt.Printf("error saving user. %v", result.Error.Error())
	}
	fmt.Printf("saved user %v.", user.Id)

	return nil
}

func existEmail(email string) (*bool,error) {
	var us = &User{Email: email}
	userGorm := config.DbGorm.First(&us);
	if userGorm.Error != nil{
		fmt.Println(userGorm.Error.Error())
		return nil, userGorm.Error
	}

	exists := userGorm.RowsAffected > 0
	return &exists, nil
}
// End Exist Email

// Init insert post
func (post *Post) insertPost () error  {
	syncSession.Lock()
	defer syncSession.Unlock()
	sql := "INSERT posts SET user_id=?, title=?,content=?"
	_, err :=config.Execute(sql,post.UserId, post.Title,post.Content)
	if err != nil {
		return err
	}
	return nil

}

// End update post

// Init insert post
//(post *Post)
func updatePostSql (postId string, userId string, title, content string) error  {
	syncSession.Lock()
	defer syncSession.Unlock()
	sql := "UPDATE posts SET title=?,content=? WHERE user_id=? and id=?"
	_, err :=config.Execute(sql,title,content,userId, postId )
	if err != nil {
		return err
	}
	return nil

}

// End update post


func deletePost (postId string)  error {
	/*syncSession.Lock()
	defer syncSession.Unlock()*/
	sql := "DELETE FROM posts WHERE id=?"
	fmt.Println(postId)
	 _  , err := config.Execute(sql,postId)
	if err != nil {
		return errors.New("Error deleting post")
	}
//	fmt.Println("el Resultado de execute: ", result)
//	fmt.Println("el error de execute: ", err)
	 return nil

}


