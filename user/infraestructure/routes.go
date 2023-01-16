package user_infra

import (
	"net/http"
	"golang.com/forum/routes"
	user_controllers "golang.com/forum/user/infraestructure/controllers"

)


func GetRoutes() *[]routes.Routes {
	return &[]routes.Routes{
		{
			Path: "/sign-up/", 
			Name: "register", 
			Methods: []string{http.MethodPost, http.MethodOptions}, 
			Handler: user_controllers.SingUp, 
			NeedsAuth: false,
		},
		{
			Path: "/login/", 
			Name: "login", 
			Methods: []string{http.MethodPost, http.MethodOptions}, 
			Handler: user_controllers.SingIn, 
			NeedsAuth: false,
		},
		{
			Path: "/", 
			Name: "test-check-out", 
			Methods: []string{http.MethodGet}, 
			Handler: user_controllers.Index, NeedsAuth: true,
	    },
	}

}