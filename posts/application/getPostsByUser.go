package post_application

import (
	"forum/config"
	"forum/posts/response"
	postUtils "forum/posts/utils"
)

func GetPostsByUser(userId string, pagination config.Pagination) (*response.GetPostsDtoPaginated, error) {
	posts, err := postUtils.GetPostsByUser(userId, pagination)
	if err != nil {
		return nil, err
	}
	postsDto, _ := response.Transform(posts)
	return postsDto, nil

}
