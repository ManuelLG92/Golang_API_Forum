package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

// Init get id by URL
func getIdByUrl(r *http.Request)  string {
	vars := mux.Vars(r)
	return vars["id"]
}

func CheckPassword(current string, hash string) error  {
	return bcrypt.CompareHashAndPassword([]byte(current),[]byte(hash))
}
