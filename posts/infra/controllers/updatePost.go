package post_infra_controllers

import (
	"fmt"
	"net/http"

	"forum/auth"
	"forum/handlers"
	"forum/helpers"
	post_application "forum/posts/application"
	post_domain "forum/posts/domain"
)

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	postId := handlers.GetFieldByUrl(r, "id")
	if postId == "" {
		helpers.SendUnprocessableEntity(w, "<id> parameter is required.")
		return
	}
	var userId string = *auth.GetUserIdFromContext(r.Context())
	post, err := helpers.DecodeBody[post_domain.PostUpdatableFields](r.Body, "Error trying to update the post")
	if err != nil {
		helpers.SendUnprocessableEntity(w, err.Error())
		return
	}

	response, err := post_application.UpdatePost(userId, postId, *post)
	if err != nil {
		fmt.Printf("Error trying to Update post: %v. Error: %v", postId, err.Error())
		fmt.Println("Error trying to Update post: ", postId)
		helpers.SendNotFound(w, fmt.Sprintf("Post %v not found", postId))
		return
	}
	fmt.Println("Post Updated")
	helpers.SendData(w, &response)
}
