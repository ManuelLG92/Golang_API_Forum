package post_application

import (
	"fmt"
	postUtils "forum/posts/utils"
)

func DeletePost(userId string, postId string) error {

	post, err := postUtils.GetPostByIdAndUserId(postId, userId)
	if err != nil {
		fmt.Printf("Error on delete. %v", err.Error())
		return err
	}
	if err := postUtils.Delete(post.Id); err != nil {
		fmt.Printf("Error on delete. %v", err.Error())
		return err
	}
	return nil

}
