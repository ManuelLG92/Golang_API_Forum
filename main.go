package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"golang.com/forum/config"
	post_infra_routes "golang.com/forum/posts/infra/routes"
	"golang.com/forum/routes"
	user_infra "golang.com/forum/user/infraestructure"
)
func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}
func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
      w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}
func main() {
	muxRouter := mux.NewRouter()
	port := ":2000"
	config.Connection()
	defer config.CloseGormConnection()

	var mapRoutes []routes.Routes
	userRoutes := user_infra.GetRoutes();
	postRoutes := post_infra_routes.GetRoutes()
	routes.AutoMigrate()
	mapRoutes = append(mapRoutes, *postRoutes...)
	mapRoutes = append(mapRoutes, *userRoutes...)

	enableCORS(muxRouter)
	err := routes.Register(mapRoutes, muxRouter)
	if err != nil {
		fmt.Printf("Errors %v", err)
		panic(fmt.Sprintf("Has been an error registerin the routes. Message: %v", err))
	}
	

	log.Println("El servidor esta a la escucha en el puerto ", port)
	log.Fatal(http.ListenAndServe(port, muxRouter))

}
