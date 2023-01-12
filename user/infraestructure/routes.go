package user_infra

import (
	"net/http"
	"golang.com/forum/routes"
)


func GetRoutes() *[]routes.Routes {
	var signUp routes.Routes = routes.Routes{Path: "/sign-up/", Name: "register", Methods: []string{http.MethodPost, "OPTIONS"}, Handler: SingUp, NeedsAuth: false}
	var signIn routes.Routes = routes.Routes{Path: "/login/", Name: "login", Methods: []string{http.MethodPost, "OPTIONS"}, Handler: SingIn, NeedsAuth: false}
	var index routes.Routes = routes.Routes{Path: "/", Name: "login", Methods: []string{http.MethodGet}, Handler: Index, NeedsAuth: true}

	var UserRoutes []routes.Routes = []routes.Routes{}
	UserRoutes = append(UserRoutes, signUp)
	UserRoutes = append(UserRoutes, signIn)
	UserRoutes = append(UserRoutes, index)

	return &[]routes.Routes{
		{
			Path: "/sign-up/", 
			Name: "register", 
			Methods: []string{http.MethodPost, "OPTIONS"}, 
			Handler: SingUp, 
			NeedsAuth: false,
		},
		{
			Path: "/login/", 
			Name: "login", 
			Methods: []string{http.MethodPost, "OPTIONS"}, 
			Handler: SingIn, 
			NeedsAuth: false,
		},
		{
			Path: "/", 
			Name: "login", 
			Methods: []string{http.MethodGet}, 
			Handler: Index, NeedsAuth: true,
	    },
	}

}