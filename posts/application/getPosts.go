package post_application

import (
	post_domain "golang.com/forum/posts/domain"
	post_utils "golang.com/forum/posts/utils"

)
func GetPosts() (*[]post_domain.Post, error) {
	posts,err := post_utils.GetPosts()
	if  err != nil {
		return nil, err
	}
	return posts, nil

}





