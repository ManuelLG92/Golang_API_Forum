package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.com/forum/auth"
	"golang.com/forum/config"
	"golang.com/forum/handlers"
)


func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	port := ":2000"
	config.CreateConnection()
	config.CreateTables()
	config.Ping()
	defer config.CloseConnection()
	config.Connection()
	handlers.AutoMigrate()
	defer config.CloseGormConnection()

	//mux := http.NewServeMux()

	/*authHandler := auth.AuthenticatedUser(handlers.GetPostsByUser)
	muxRouter.Handle("/users/posts/{id:[0-9]+}",authHandler).Methods("GET")*/
	authHandler := auth.AuthenticatedUser(handlers.Index)
	//next := http.HandlerFunc(handlers.Index)
	muxRouter.Handle("/", authHandler).Methods("GET")
	muxRouter.HandleFunc("/users/sign-up/",handlers.SingUp).Methods("POST", "OPTIONS")
	muxRouter.HandleFunc("/users/sign-in/",handlers.SingIn).Methods("POST", "OPTIONS")
	muxRouter.HandleFunc("/login/",handlers.SingIn).Methods("POST", "OPTIONS")

	log.Println("El servidor esta a la escucha en el puerto ", port)
	log.Fatal(http.ListenAndServe(port, muxRouter))
	//muxRouter.HandleFunc("/auth/",auth.IsUserAuth).Methods("GET")


	//muxRouter.HandleFunc("/users/",handlers.GetUsers).Methods("GET")
	//muxRouter.HandleFunc("/users/{id:[0-9]+}",handlers.GetUserOnEmail).Methods("GET")
	//muxRouter.HandleFunc("/users/",handlers.UpdateUsers).Methods("PUT")
	//muxRouter.HandleFunc("/posts/newPost/",handlers.CreatePost).Methods("POST", "OPTIONS")
	//muxRouter.HandleFunc("/users/newPost/",handlers.CreatePost).Methods("POST")



	/*muxRouter.HandleFunc("/posts/",handlers.GetPosts).Methods("GET")
	muxRouter.HandleFunc("/posts/edit/{id:[0-9]+}/",handlers.EditPost).Methods("POST","OPTIONS")
	muxRouter.HandleFunc("/post/{id:[0-9]+}/",handlers.GetPostsById).Methods("GET", "OPTIONS")
*/
	/*
	// Init auth user for edit
	authHandlerEdit := auth.AuthenticatedUser(handlers.UpdatePost)
	muxRouter.Handle("/users/posts/{id:[0-9]+}",authHandlerEdit).Methods("PUT","OPTIONS")*/
	//muxRouter.HandleFunc("/post/edit/{id:[0-9]+}/",handlers.UpdatePost).Methods("PUT","OPTIONS")
	// End auth user for edit

	//muxRouter.HandleFunc("/post/delete/{id:[0-9]+}/",handlers.DeletePost).Methods("POST","OPTIONS")

	//Init Auth users test
	/*authHandler := auth.AuthenticatedUser(handlers.GetPostsByUser)
	muxRouter.Handle("/users/posts/{id:[0-9]+}",authHandler).Methods("GET")*/
	//End Auth users test
/*
	authRouter := muxRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/users",handlers.GetUsers).Methods("GET")
	authRouter.HandleFunc("/users/{id:[0-9]+}",handlers.GetUserByEmail).Methods("GET")
	authRouter.HandleFunc("/users/forum",handlers.GetPosts).Methods("GET")
	authRouter.HandleFunc("/users/{id:[0-9]+}/posts",handlers.GetPostsByUser).Methods("GET")*/
	//Auth routes

	// Trying concurrency, channels and context.
	/*
	log := log.New(os.Stdout, "product-api", log.LstdFlags)
	server := &http.Server{
		Addr: port,
		Handler: muxRouter,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,

	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	signalChanel := make(chan os.Signal)
	signal.Notify(signalChanel,os.Interrupt)
	signal.Notify(signalChanel, os.Kill)

	signal := <- signalChanel
	log.Println("received terminated, graceful shutdown", signal)

	tc,_ := context.WithTimeout(context.Background(),30*time.Second)
	server.Shutdown(tc)*/
	// All new from Nic
/* 	server := &http.Server{
		Addr: port,
		Handler: muxRouter,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,

	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	signalChanel := make(chan os.Signal)
	signal.Notify(signalChanel,os.Interrupt)
	signal.Notify(signalChanel, os.Kill)

	signal := <- signalChanel
	log.Println("received terminated, graceful shutdown", signal)

	tc,_ := context.WithTimeout(context.Background(),30*time.Second)
	server.Shutdown(tc) */

	

}

