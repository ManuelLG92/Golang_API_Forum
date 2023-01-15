package post_infra_routes

import (
	"net/http"
	"golang.com/forum/routes"
	post_infra_controllers "golang.com/forum/posts/infra/controllers"

)

func GetRoutes() *[]routes.Routes {

	return &[]routes.Routes{
		{
			Path: "/post", 
			Name: "post-create", 
			Methods: []string{http.MethodPost, http.MethodOptions}, 
			Handler: post_infra_controllers.CreatePost, 
			NeedsAuth: true,
		},
		{
			Path: "/post", 
			Name: "post-delete", 
			Methods: []string{http.MethodDelete, http.MethodOptions}, 
			Handler: post_infra_controllers.DeletePost, 
			NeedsAuth: true,
		},
		{
			Path: "/post", 
			Name: "post-update", 
			Methods: []string{http.MethodDelete, http.MethodPut}, 
			Handler: post_infra_controllers.UpdatePost, 
			NeedsAuth: true,
		},
		{
			Path: "/posts/lastest", 
			Name: "post-latest", 
			Methods: []string{http.MethodGet}, 
			Handler: post_infra_controllers.GetPosts, 
			NeedsAuth: true,
	    },
		{
			Path: "/posts/me", 
			Name: "post-me", 
			Methods: []string{http.MethodGet}, 
			Handler: post_infra_controllers.GetPostsByUser,
			 NeedsAuth: true,
	    },
	}

}