package handlers

import (
	"../config"
	"../models"
	"fmt"
	"github.com/pkg/errors"
)

type SessionsData struct {
	Uuid string `json:"uuid"`
	UserId int `json:"user_id"`
}

// Init GetUserByEmail
func GetUserByEmail(email string) *User {
	user := newUser("","","","")
	// si lo encuentra por el email lo devuelve
	sql := "SELECT id, name,last_name, password, email FROM users where email=?"
	rows, err := config.Query(sql,email)
	if err != nil {
		return nil
		//fmt.Println("inside if", err)
	}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.LastName,&user.Password, &user.Email)
		if  err != nil {
			return nil
		}
	}
	return user
}
// End GetUserByEmail

// Init insertUser in DB
func (user *User) insertUser() error {
	syncSession.Lock()
	defer syncSession.Unlock()
	sql := "INSERT users SET name=? , last_name=?, password=?, email=?"
	Result, err := config.Execute(sql,user.Name,user.LastName , user.Password, user.Email)
	if err != nil {
		return err
	}
	fmt.Println(Result)

	return nil
}
// End insertUser in DB

//Init Exist Email
func existEmail(email string) error {
	syncSession.Lock()
	defer syncSession.Unlock()
	sql := "SELECT name, email FROM users where email=?"
	rows, _ := config.Query(sql,email)
	for rows.Next() {
		return models.ErrorUserRegistred
		//rows.Scan(&user.Password, &user.Email)
	}
	return nil
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
func updatePostSql (postId , userId int, title, content string) error  {
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

func getSessionByUuid (uuid string) (error, int)  {
	//syncSession.Lock()
	//defer syncSession.Unlock()
	var userID int
	sql := "SELECT user_id FROM sessions WHERE uuid=?"
	result , err :=config.Query(sql,uuid)
	if err != nil {
		fmt.Println("Could not catch any row")
		return err,0
	} else {
		for result.Next(){
			result.Scan(&userID)
		}
		//fmt.Println("user id in query : ", userID)
		return nil, userID
	}
}

func deletePost (postId int)  error {
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


