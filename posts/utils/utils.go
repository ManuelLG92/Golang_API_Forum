package post_utils

import (
	"errors"
	"fmt"
	"golang.com/forum/config"
	post_domain "golang.com/forum/posts/domain"
)

func AutoMigrate()  {
	err := config.Connection().AutoMigrate(&post_domain.Post{})
	if err != nil {
		return 
	}
}

func GetPostById(postId string) (*post_domain.Post, error) {
	post := &post_domain.Post{Id: postId}
	postGorm := config.DbGorm.First(&post);
	if postGorm.Error != nil{
		return nil, postGorm.Error
	}
	if post.Id == ""{
		return nil, errors.New(fmt.Sprintf("Post %v not found", postId))
	}
	return post, nil
}

func GetPostByIdAndUserId(postId , userId string) (*post_domain.Post, error) {
	fmt.Println("GetPostByIdAndUserId init",postId, userId)

	var post = &post_domain.Post{}
	if err := config.DbGorm.Where("id = ? AND user_id = ?",postId,userId).First(&post); err.Error != nil{
		fmt.Println("error quering", err)
		return nil, err.Error
	}
	fmt.Println("GetPostByIdAndUserId after where",post)

	if post.Id == ""{
		return nil, errors.New(fmt.Sprintf("Post %v not found", postId))
	}
	return post, nil
}

func GetPostsByUser(userId string) (*[]post_domain.Post, error) {
	var posts []post_domain.Post
	postGorm := config.DbGorm.Where("user_id = ?",userId).Find(&posts);
	if postGorm.Error != nil{
		return nil, postGorm.Error
	}
	return &posts, nil
}

func GetPosts() (*[]post_domain.Post, error) {
	var posts []post_domain.Post
	postGorm := config.DbGorm.Find(&posts);
	if postGorm.Error != nil{
		return nil, postGorm.Error
	}
	return &posts, nil
}


func Save(post *post_domain.Post) error {
	result := config.DbGorm.Save(&post);
	if result.Error != nil {
		fmt.Printf("error saving post. %v", result.Error.Error())
		return result.Error
	}
	fmt.Printf("saved posts %v.", post.Id)
	return nil
}

func Delete(id string) error {
	fmt.Println("init post_utils.Delete")
	
	result := config.DbGorm.Where("id = ?",id).Delete(&post_domain.Post{});
	fmt.Println("after where init post_utils.Delete")

	if result.Error != nil {
	fmt.Println("after where inside error != nil post_utils.Delete")

		fmt.Printf("error deleting post. %v", result.Error.Error())
		return result.Error
	}
	fmt.Println("after where no error post_utils.Delete")

	fmt.Printf("Deleted post %v.", id)
	return nil
}
