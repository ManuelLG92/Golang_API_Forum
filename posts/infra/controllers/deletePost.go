package post_infra_controllers

import (
	"fmt"
	"forum/auth"
	"forum/handlers"
	"forum/helpers"
	post_application "forum/posts/application"
	"net/http"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {
	postId := handlers.GetFieldByUrl(r, "id")
	if postId == "" {
		helpers.SendUnprocessableEntity(w, "<id> parameter is required.")
		return
	}
	var userId = *auth.GetUserIdFromContext(r.Context())

	fmt.Println("Before call post_application.DeletePost")
	if err := post_application.DeletePost(userId, postId); err != nil {
		fmt.Printf("Error trying to Delete post: %v. Error: %v", postId, err.Error())
		fmt.Println("Error trying to Delete post: ", postId)
		helpers.SendNotFound(w, fmt.Sprintf("Post %v not found", postId))
		return
	}
	fmt.Println("Post Deleted")
	helpers.SendNoContent(w)

}
