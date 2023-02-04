package post_application

import (
	"fmt"
	"forum/posts/response"
	postUtils "forum/posts/utils"
)

type PostWithUserInformation struct {
	Id           string `json:"id"`
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	UserLastName string `json:"user_last_name"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func GetPosts() (*[]response.GetPostsDto, error) {
	posts, err := postUtils.GetPosts()
	if err != nil {
		fmt.Println("Error -> postUtils.GetPosts ", err.Error())
		return nil, err
	}
	postsDto, err := response.Transform(posts)
	return postsDto, nil

}
