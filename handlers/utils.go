package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/constraints"
)

func getIdByUrl(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["id"]
}

func CheckPassword(current string, hash string) error  {
	return bcrypt.CompareHashAndPassword([]byte(current),[]byte(hash))
}

type Primitives interface {
	string|bool|constraints.Integer|constraints.Complex
 }
func Equals[T Primitives](left T, right T) bool {
	return left == right;
}
