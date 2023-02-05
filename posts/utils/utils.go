package post_utils

import (
	"errors"
	"fmt"
	"forum/config"
	postDomain "forum/posts/domain"
	"forum/storage"
)

var cache = storage.NewCache[postDomain.Post]()

func GetPostByIdAndUserId(postId, userId string) (*postDomain.Post, error) {
	var composedCacheKey = postId + "-" + userId

	if result := cache.Get(composedCacheKey); result != nil {
		return result, nil
	}
	fmt.Println("GetPostByIdAndUserId init", postId, userId)

	var post = &postDomain.Post{}
	if err := config.DbGorm.Where("id = ? AND user_id = ?", postId, userId).First(&post); err.Error != nil {
		fmt.Println("error quering", err)
		return nil, err.Error
	}
	fmt.Println("GetPostByIdAndUserId after where", post)

	if post.Id == "" {
		return nil, errors.New(fmt.Sprintf("Post %v not found", postId))
	}
	cache.Set(postId, *post)
	return post, nil
}

func GetPostsByUser(userId string) (*[]postDomain.Post, error) {

	var posts []postDomain.Post
	postGorm := config.DbGorm.Where("user_id = ?", userId).Find(&posts)
	if postGorm.Error != nil {
		return nil, postGorm.Error
	}
	return &posts, nil
}

func GetPosts() (*[]postDomain.Post, error) {
	var posts []postDomain.Post
	postGorm := config.DbGorm.Find(&posts)
	if postGorm.Error != nil {
		return nil, postGorm.Error
	}
	return &posts, nil
}

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
