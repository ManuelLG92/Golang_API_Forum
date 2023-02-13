package main

import (
	"context"
	"fmt"
	"forum/config"
	"forum/helpers"
	postInfraRoutes "forum/posts/infra/routes"
	"forum/routes"
	userInfra "forum/user/infraestructure"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
)

const ContextStartAt = "start-at"
const PaginateFromMiddleware = "paginate"

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}
func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			startReq := time.Now()
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			//w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-access-token. X-Access-Token")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Expose-Headers", "x-access-token, X-Access-Token")
			fmt.Println("req body start request", req.Body)
			ctx := req.Context()
			ctx = context.WithValue(ctx, ContextStartAt, startReq)
			if req.Method == http.MethodGet {
				var paginate config.Pagination
				if validations := paginate.PaginateFromUrlQueryParams(req.URL.Query()); validations != nil {
					helpers.SendUnprocessableEntity(w, strings.Join(*validations, ","))
					return
				}
				ctx = context.WithValue(ctx, PaginateFromMiddleware, paginate)
			}
			defer fmt.Printf("Request time: %v ", time.Since(startReq).Microseconds())
			next.ServeHTTP(w, req.WithContext(ctx))
			//next.ServeHTTP(w, req)
		})
}
func main() {
	helpers.LoadEnvs()
	router := mux.NewRouter()
	port := ":2000"
	config.Connection()
	defer config.CloseGormConnection()

	var mapRoutes []routes.Routes
	userRoutes := userInfra.GetRoutes()
	postRoutes := postInfraRoutes.GetRoutes()
	routes.AutoMigrate()
	mapRoutes = append(mapRoutes, *postRoutes...)
	mapRoutes = append(mapRoutes, *userRoutes...)

	enableCORS(router)
	err := routes.Register(mapRoutes, router)
	if err != nil {
		fmt.Printf("Errors %v", err)
		panic(fmt.Sprintf("Has been an error registering the routes. Message: %v", err))
	}

	log.Println("Server listening on port  #", port)
	log.Fatal(http.ListenAndServe(port, router))

}
