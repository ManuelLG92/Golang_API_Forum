package post_application

import (
	post_domain "golang.com/forum/posts/domain"
	post_utils "golang.com/forum/posts/utils"
)

func CreatePost(userId string, data post_domain.PostUpdatableFields) (*post_domain.Post, error) {
	post, errorValidPost := post_domain.NewPost(userId, data.Title, data.Content)
	if errorValidPost != nil {
		return nil, errorValidPost
	}
	if err := post_utils.Save(post); err != nil {
		return nil, err
	}
	return post, nil

}
