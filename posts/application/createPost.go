package post_application

import (
	postDomain "forum/posts/domain"
	postUtils "forum/posts/utils"
)

func CreatePost(userId string, data postDomain.PostUpdatableFields) (*postDomain.Post, error) {
	post, errorValidPost := postDomain.NewPost(userId, data.Title, data.Content)
	if errorValidPost != nil {
		return nil, errorValidPost
	}
	if err := postUtils.Save(post); err != nil {
		return nil, err
	}
	return post, nil

}
