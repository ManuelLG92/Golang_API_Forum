package post_infra_controllers

import (
	"fmt"
	"forum/auth"
	"forum/config"
	"forum/helpers"
	post_application "forum/posts/application"
	"net/http"
)

func GetPostsByUser(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())
	fmt.Printf("Posts fetched by user %v.", userId)

	paginate, errCtx := r.Context().Value("paginate").(config.Pagination)
	if !errCtx {
		fmt.Printf("no pagination found")
		helpers.SendUnprocessableEntity(w, "no pagination found")
	}

	posts, err := post_application.GetPostsByUser(*userId, paginate)
	if err != nil {
		fmt.Println("Error trying to get the posts: ", err)
		helpers.SendInternalServerError(w)
		return
	}
	helpers.SendData(w, &posts)

}
