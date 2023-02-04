package routes

import (
	"errors"
	"fmt"
	"forum/auth"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slices"
	"net/http"
)

const (
	GET     string = http.MethodGet
	POST    string = http.MethodPost
	PUT     string = http.MethodPut
	PATCH   string = http.MethodPatch
	DELETE  string = http.MethodDelete
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
	var routesErrors []error
	count := 0

	for _, route := range routes {
		if err := route.validate(); err != nil {
			routesErrors = append(routesErrors, err)
			continue
		}
		if len(routesErrors) == 0 {

			fmt.Println("router methods", route.Methods)
			if route.NeedsAuth {
				router.Handle(route.Path,
					auth.StartRequest(auth.AuthenticatedUser(http.HandlerFunc(route.Handler)))).Methods(route.Methods...).Name(route.Name)
			} else {
				router.Handle(route.Path,
					auth.StartRequest(http.HandlerFunc(route.Handler))).Methods(route.Methods...).Name(route.Name)
			}
			count++
			fmt.Printf("%v. Route: %v - Protected?: %v - Method/s: %v", count, route.Path, route.NeedsAuth, route.Methods)
			fmt.Println()

		}
	}
	if len(routesErrors) == 0 {

		fmt.Printf("Number of routes: %v", count)
		fmt.Println("")
		return nil
	}
	return routesErrors

}
