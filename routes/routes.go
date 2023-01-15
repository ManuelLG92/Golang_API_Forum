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
	GET     string = http.MethodGet
	POST    string = http.MethodPost
	PUT     string = http.MethodPut
	PATCH   string = http.MethodPatch
	DELETE     string = http.MethodDelete
	OPTIONS string = http.MethodOptions
)

type Routes struct {
	Path    string
	Name    string
	Methods []string
	Handler func(http.ResponseWriter,
		*http.Request)
	NeedsAuth bool
}

var availableMethods = []string{GET, POST, PUT, PATCH, DELETE, OPTIONS}

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
	count := 0

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
			count++
			fmt.Printf("%v. Route: %v - Protected?: %v - Method: %v", count,route.Path, route.NeedsAuth, route.Methods)
			fmt.Println()
			
		}
	}
	if len(errors) == 0 {
		fmt.Printf("Number of routes: %v", count)
	    fmt.Println("")
		return nil
	}

	return errors

}
