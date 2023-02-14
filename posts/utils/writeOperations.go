package post_utils

import (
	"fmt"
	"forum/config"
	postDomain "forum/posts/domain"
)

func Save(post *postDomain.Post) error {
	result := config.DbGorm.Save(&post)
	if result.Error != nil {
		fmt.Printf("error saving post. %v", result.Error.Error())
		return result.Error
	}
	fmt.Printf("saved posts %v.", post.Id)
	return nil
}

func Delete(id string) error {
	fmt.Println("init post_utils.Delete")

	result := config.DbGorm.Where("id = ?", id).Delete(&postDomain.Post{})
	fmt.Println("after where init post_utils.Delete")

	if result.Error != nil {
		fmt.Println("after where inside error != nil post_utils.Delete")

		fmt.Printf("error deleting post. %v", result.Error.Error())
		return result.Error
	}
	fmt.Println("after where no error post_utils.Delete")

	fmt.Printf("Deleted post %v.", id)

	cache.Delete(id)
	return nil
}
