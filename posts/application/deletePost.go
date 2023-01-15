package post_application

import (
	post_utils "golang.com/forum/posts/utils"
)


func DeletePost(userId string, postId string) error {
	post, err := post_utils.GetPostByIdAndUserId(postId, userId)
	if err != nil {
		return err
	}
	if err := post_utils.Delete(post); err != nil {
		return err
	}
	return nil

}
