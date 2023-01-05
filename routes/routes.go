package routes

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"golang.com/forum/auth"
	"golang.org/x/exp/slices"
	"net/http"
)

const (
	GET     string = "GET"
	POST    string = "POST"
	PUT     string = "PUT"
	PATCH   string = "PATCH"
	OPTIONS string = "OPTIONS"
)

type Routes struct {
	Path    string
	Name    string
	Methods []string
	Handler func(http.ResponseWriter,
		*http.Request)
	NeedsAuth bool
}

var availableMethods = []string{GET, POST, PUT, PATCH, OPTIONS}

func (r *Routes) validate() error {
	for _, v := range r.Methods {
		if slices.Contains(availableMethods, v) == false {
			return errors.New(fmt.Sprintf("Method %v not allowed", v))
		}
	}
	return nil
}

func Register(routes []Routes, router *mux.Router) []error {
	var errors []error
	fmt.Printf("count router %v", len(routes))
	fmt.Println("")

	for _, route := range routes {
		if err := route.validate(); err != nil {
			errors = append(errors, err)
		}
		if len(errors) == 0 {

			if route.NeedsAuth {
				router.Handle(route.Path,
					auth.StartRequest(auth.AuthenticatedUser(http.HandlerFunc(route.Handler)))).Methods(route.Methods...)
			} else {
				router.Handle(route.Path,
					auth.StartRequest(http.HandlerFunc(route.Handler))).Methods(route.Methods...)
			}
			fmt.Printf(" route path %v", route.Path)
			fmt.Println("")
			fmt.Printf(" route methods %v", route.Methods)
			fmt.Println("end")
		}
	}
	if len(errors) == 0 {
		return nil
	}

	return errors

}
