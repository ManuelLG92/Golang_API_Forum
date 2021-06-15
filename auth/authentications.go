package auth

import (
	"encoding/json"
	"golang.com/forum/handlers"
	"net/http"
)


type CustomHandler func(w http.ResponseWriter, r *http.Request)

func AuthenticatedUser (function CustomHandler ) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	isAuth := handlers.IsAuthenticated(r)
	if isAuth == nil{
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode("You need to be logged in to see this content. ")
		return
	}
	function(w,r)
	})
}

func IsUserAuth(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	isAuth := handlers.IsAuthenticated(r)
	if isAuth == nil {
		w.WriteHeader(http.StatusUnauthorized)
		_= json.NewEncoder(w).Encode("You need be authenticated to see this content. ")
		return
	} else {
		w.WriteHeader(http.StatusAccepted)
		_=json.NewEncoder(w).Encode(isAuth)
		return
	}
}


