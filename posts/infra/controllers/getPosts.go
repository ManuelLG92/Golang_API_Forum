package post_infra_controllers

import (
	"fmt"
	"forum/auth"
	"forum/config"
	"forum/helpers"
	postApplication "forum/posts/application"
	"net/http"
	"strings"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())
	fmt.Printf("Posts fetched by user %v.", userId)

	var paginate config.Pagination
	query := r.URL.Query()
	if validations := paginate.PaginateFromUrlQueryParams(query); validations != nil {
		helpers.SendUnprocessableEntity(w, strings.Join(*validations, ","))
		return
	}
	posts, err := postApplication.GetPosts(paginate)
	if err != nil {
		fmt.Println("Error trying to get the posts: ", err)
		helpers.SendInternalServerError(w)
		return
	}
	helpers.SendData(w, &posts)

}
