package post_infra_controllers

import (
	"fmt"
	"net/http"

	"golang.com/forum/auth"
	"golang.com/forum/handlers"
	"golang.com/forum/helpers"
	post_application "golang.com/forum/posts/application"
	post_domain "golang.com/forum/posts/domain"
)

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	postId := handlers.GetFieldByUrl(r, "id")
	if postId == "" {
		http.Error(w, "No valid post Id", http.StatusNotAcceptable)
		return
	}
	var userId string = *auth.GetUserIdFromContext(r.Context())
	post, err := helpers.DecodeBody[post_domain.PostUpdatableFields](r.Body, "Error tryig to update the post. ")
	if err != nil {
		helpers.SendUnprocessableEntity(w, err.Error())
	}

	response, err := post_application.UpdatePost(userId, postId, *post)
	if err != nil {
		fmt.Printf("Error trying to Update post: %v. Error: %v", postId, err.Error())
		fmt.Println("Error trying to Update post: ", postId)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Post Updated")
	helpers.SendData(w, &response)
}
