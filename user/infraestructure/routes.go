package user_infra

import (
	"forum/routes"
	user_controllers "forum/user/infraestructure/controllers"
	"net/http"
)

func GetRoutes() *[]routes.Routes {
	return &[]routes.Routes{
		{
			Path:      "/sign-up",
			Name:      "register",
			Methods:   []string{http.MethodPost, http.MethodOptions},
			Handler:   user_controllers.SingUp,
			NeedsAuth: false,
		},
		{
			Path:      "/login",
			Name:      "login",
			Methods:   []string{http.MethodPost, http.MethodOptions},
			Handler:   user_controllers.SingIn,
			NeedsAuth: false,
		},
		{
			Path:    "/",
			Name:    "check-auth",
			Methods: []string{http.MethodGet},
			Handler: user_controllers.Index, NeedsAuth: true,
		},
	}

}
