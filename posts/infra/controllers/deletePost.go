package post_infra_controllers

import (
	"fmt"
	"golang.com/forum/auth"
	"golang.com/forum/handlers"
	"net/http"
	"golang.com/forum/helpers"
	post_application "golang.com/forum/posts/application"
)


func DeletePost(w http.ResponseWriter, r *http.Request) {
	postId := handlers.GetFieldByUrl(r, "id")
	if postId == "" {
		http.Error(w, "No valid post Id", http.StatusNotAcceptable)
		return
	}
	var userId string = *auth.GetUserIdFromContext(r.Context())

	if err := post_application.DeletePost(userId, postId); err != nil {
		fmt.Printf("Error trying to Update post: %v. Error: %v", postId, err.Error())
		fmt.Println("Error trying to Update post: ", postId)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Post Updated")
	helpers.SendNoContent(w)

}
