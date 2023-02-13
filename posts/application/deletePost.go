package post_application

import (
	"fmt"
	postUtils "forum/posts/utils"
)

func DeletePost(userId string, postId string) error {
	fmt.Println("init post_application.DeletePost")

	post, err := postUtils.GetPostByIdAndUserId(postId, userId)
	if err != nil {
		fmt.Printf("Error on delete. %v", err.Error())
		return err
	}
	fmt.Println("error result postUtils.Delete", err)
	fmt.Println("post result postUtils.Delete", post)
	fmt.Println("before postUtils.Delete")
	if err := postUtils.Delete(post.Id); err != nil {
		fmt.Printf("Error on delete. %v", err.Error())
		return err
	}
	return nil

}
