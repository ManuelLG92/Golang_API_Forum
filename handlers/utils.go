package handlers

import (
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/constraints"
	"net/http"
)

func GetFieldByUrl(r *http.Request, fieldName string) string {
	vars := mux.Vars(r)
	return vars[fieldName]
}

func CheckHashString(current string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(current), []byte(hash))
}

type Primitives interface {
	string | bool | constraints.Integer | constraints.Complex
}

func Equals[T Primitives](left T, right T) bool {
	return left == right
}
