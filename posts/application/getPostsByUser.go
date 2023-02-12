package post_application

import (
	"forum/posts/response"
	postUtils "forum/posts/utils"
)

func GetPostsByUser(userId string) (*response.GetPostsDtoPaginated, error) {
	posts, err := postUtils.GetPostsByUser(userId)
	if err != nil {
		return nil, err
	}
	postsDto, _ := response.Transform(posts)
	return postsDto, nil

}
