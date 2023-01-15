package post_application

import (
	post_domain "golang.com/forum/posts/domain"
	post_utils "golang.com/forum/posts/utils"
)

func UpdatePost(userId string, postId string, data post_domain.PostUpdatableFields) (*post_domain.Post, error) {
	postModel, err := post_utils.GetPostByIdAndUserId(postId, userId)
	if err != nil {
		return nil, err
	}
	post, err := postModel.EditPost(&data)
	if err != nil {
		return nil, err
	}
	post_utils.Save(post)
	return post, nil

}

