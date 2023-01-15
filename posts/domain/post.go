package post_domain

import (
	"time"

	"github.com/google/uuid"
	"golang.com/forum/helpers"
)

type Post struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PostUpdatableFields struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Posts []Post

func (post *Post) ValidPost() error {
	if err := post.validPostData(); err != nil {
		return err
	}
	return nil
}

func (post *Post) validPostData() error {
	if len(post.Title) < 1 || len(post.Title) > 50 || len(post.Content) < 1 || len(post.Content) > 255 {
		return helpers.ErrorPostData
	}
	return nil
}

func (post *Post) EditPost(data *PostUpdatableFields) (*Post, error) {
	post.Title = data.Title
	post.Content = data.Content
	post.UpdatedAt = time.Now().String()
	if err := post.ValidPost(); err != nil {
		return nil, err
	}
	return post, nil

}

func NewPost(userId string, title string, content string) (*Post, error) {

	post := &Post{Id: uuid.New().String(), UserId: userId, Title: title, Content: content}
	if err := post.ValidPost(); err != nil {
		return nil, err
	}
	return post, nil

}
