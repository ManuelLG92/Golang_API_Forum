package user_controllers

import (
	"fmt"
	"forum/auth"
	"forum/helpers"
	user_application "forum/user/application"
	user_domain "forum/user/domain"
	"log"
	"net/http"
	"strings"
)

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=64"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("before ctx.")
	userCtx := auth.GetUserIdFromContext(r.Context())
	fmt.Printf("user ctx. %v", *userCtx)
	_, err := fmt.Fprintf(w, "You are in golang app!")
	if err != nil {
		return
	}
}

func SingUp(w http.ResponseWriter, r *http.Request) {
	var user *user_domain.User

	user, err := helpers.DecodeBody[user_domain.User](r.Body, "missing fields on user")
	if err != nil {
		log.Printf("Error validaing user %v", err)
		helpers.SendUnprocessableEntity(w, err.Error())
		return
	}

	user, err = user_application.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user %v", err.Error())
		helpers.SendUnprocessableEntity(w, err.Error())
		return
	}
	helpers.SendCreated(w, "User created")
}

func SingIn(w http.ResponseWriter, r *http.Request) {

	var data *Login

	data, err := helpers.DecodeBody[Login](r.Body, "missing fields on user")
	fmt.Println(r)
	if err != nil {
		log.Printf("Error parsing user %v", err.Error())
		helpers.SendUnprocessableEntity(w, err.Error())
		return
	}
	data.Email = strings.TrimSpace(data.Email)
	data.Password = strings.TrimSpace(data.Password)
	token, err := user_application.Login(data.Email, data.Password, r.RemoteAddr)
	if err != nil {
		fmt.Printf("Error checking user credentials. %v", err.Error())
		helpers.SendCustom(w, err.Error(), 400)
		return
	}
	fmt.Println("authenticated user")
	w.Header().Set("x-access-token", *token)
	w.WriteHeader(http.StatusOK)
}
