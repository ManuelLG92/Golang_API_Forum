package routes

import (
	"forum/config"
	post_domain "forum/posts/domain"
	user_domain "forum/user/domain"
)

func AutoMigrate() {
	err := config.Connection().AutoMigrate(&user_domain.User{}, &post_domain.Post{})
	if err != nil {
		return
	}
}
