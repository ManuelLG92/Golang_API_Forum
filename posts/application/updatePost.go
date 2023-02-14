package post_application

import (
	"fmt"
	postDomain "forum/posts/domain"
	postUtils "forum/posts/utils"
)

func UpdatePost(userId string, postId string, data postDomain.PostUpdatableFields) (*postDomain.Post, error) {
	postModel, err := postUtils.GetPostByIdAndUserId(postId, userId)
	if err != nil {
		return nil, err
	}
	post, err := postModel.EditPost(&data)
	if err != nil {
		return nil, err
	}
	if err := postUtils.Save(post); err != nil {
		fmt.Printf("Error updating the post. %v", err.Error())
		return nil, err
	}
	return post, nil

}
