package post_application

import (
	"fmt"

	post_utils "golang.com/forum/posts/utils"
)


func DeletePost(userId string, postId string) error {
	fmt.Println("init post_application.DeletePost")

	post, err := post_utils.GetPostByIdAndUserId(postId, userId)
	if err != nil {
		fmt.Printf("Error on delete. %v", err.Error())
		return err
	}
	fmt.Println("error result post_utils.Delete", err)
	fmt.Println("post result post_utils.Delete", post)
	fmt.Println("before post_utils.Delete")
	if err := post_utils.Delete(post.Id); err != nil {
		fmt.Printf("Error on delete. %v", err.Error())
		return err
	}
	return nil

}
