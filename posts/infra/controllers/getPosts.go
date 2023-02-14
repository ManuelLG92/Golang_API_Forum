package post_infra_controllers

import (
	"errors"
	"fmt"
	"forum/auth"
	"forum/config"
	"forum/helpers"
	postApplication "forum/posts/application"
	postUtils "forum/posts/utils"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())
	fmt.Printf("Posts fetched by user %v.", userId)

	paginate, errCtx := r.Context().Value("paginate").(config.Pagination)
	if !errCtx {
		fmt.Printf("no pagination found")
		helpers.SendUnprocessableEntity(w, "no pagination found")
	}

	posts, err := postApplication.GetPosts(paginate)
	if err != nil {
		var badReqError postUtils.BadRequestError
		if errors.As(err, &badReqError.Err) {
			fmt.Println(err.Error())
			helpers.SendBadRequestEntity(w, err.Error())
			return
		}
		fmt.Println("Error trying to get the posts: ", err)
		helpers.SendInternalServerError(w)
		return
	}
	helpers.SendData(w, &posts)

}
