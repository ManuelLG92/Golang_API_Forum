package post_infra_controllers

import (
	"fmt"
	"net/http"

	"golang.com/forum/auth"
	"golang.com/forum/helpers"
	post_application "golang.com/forum/posts/application"
	post_domain "golang.com/forum/posts/domain"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())
	post, err := helpers.DecodeBody[post_domain.PostUpdatableFields](r.Body, "Error")
	if err != nil {
		helpers.SendUnprocessableEntity(w, fmt.Sprintf("Error trying fit the body to struct. %v", err.Error()))
		return
	}

	createdPost, err := post_application.CreatePost(*userId, *post)
	if err != nil {
		fmt.Println("Error trying to create a user: ", err)
		helpers.SendNoContent(w)
		return
	}
	fmt.Println("Post created")
	helpers.SendCreated(w, createdPost.Id)

}
