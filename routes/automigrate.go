package routes

import (
	"golang.com/forum/config"
	user_domain "golang.com/forum/user/domain"
	post_domain "golang.com/forum/posts/domain"

)

func AutoMigrate()  {
	err := config.Connection().AutoMigrate(&user_domain.User{},&post_domain.Post{})
	if err != nil {
		return 
	}
}