package response

import (
	"fmt"
	postDomain "forum/posts/domain"
	"forum/user/infraestructure/entryPoint"
)

type GetPostsDto struct {
	Id           string `json:"id"`
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	UserLastName string `json:"user_last_name"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func Transform(posts *[]postDomain.Post) (*[]GetPostsDto, error) {

	var errors []string
	var postsDto []GetPostsDto
	for _, post := range *posts {
		user, err := entryPoint.FetchUserById(post.UserId)
		if err != nil {
			errors = append(errors, err.Error())
			continue
		}

		dto := GetPostsDto{
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
	if len(errors) > 0 {
		fmt.Printf("Errors -> %v \n", errors)
	}
	return &postsDto, nil

}
