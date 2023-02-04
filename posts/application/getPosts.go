package post_application

import (
	"fmt"
	postUtils "forum/posts/utils"
	"forum/user/infraestructure/entryPoint"
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

func GetPosts() (*[]PostWithUserInformation, error) {
	posts, err := postUtils.GetPosts()
	if err != nil {
		fmt.Println("Error -> postUtils.GetPosts ", err.Error())
		return nil, err
	}
	var errors []string
	var postsDto []PostWithUserInformation
	for _, post := range *posts {
		user, err := entryPoint.FetchUserById(post.UserId)
		if err != nil {
			errors = append(errors, err.Error())
			continue
		}

		dto := PostWithUserInformation{
			Id:           post.Id,
			UserId:       post.UserId,
			UserName:     user.Name,
			UserLastName: user.LastName,
			Title:        post.Title,
			Content:      post.Content,
			CreatedAt:    post.CreatedAt,
			UpdatedAt:    post.UpdatedAt}
		postsDto = append(postsDto, dto)

	}
	fmt.Printf("Errors -> %v \n", errors)
	return &postsDto, nil

}
