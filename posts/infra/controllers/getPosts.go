package post_infra_controllers

import (
	"fmt"
	"forum/auth"
	"forum/helpers"
	post_application "forum/posts/application"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())
	fmt.Printf("Posts fetched by user %v.", userId)

	posts, err := post_application.GetPosts()
	if err != nil {
		fmt.Println("Error trying to get the posts: ", err)
		helpers.SendInternalServerError(w)
		return
	}
	helpers.SendData(w, &posts)

}
