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

	err := routes.Register(mapRoutes, muxRouter)
	if err != nil {
		fmt.Printf("Errors %v", err)
		panic(fmt.Sprintf("Has been an error registerin the routes. Message: %v", err))
	}

	log.Println("El servidor esta a la escucha en el puerto ", port)
	log.Fatal(http.ListenAndServe(port, muxRouter))

}
