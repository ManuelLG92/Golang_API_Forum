package post_application

import (
	post_domain "forum/posts/domain"
	post_utils "forum/posts/utils"
)

func GetPostsByUser(userId string) (*[]post_domain.Post, error) {
	posts, err := post_utils.GetPostsByUser(userId)
	if err != nil {
		return nil, err
	}
	return posts, nil

}
