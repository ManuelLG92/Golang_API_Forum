package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type CustomHandler func(w http.ResponseWriter, r *http.Request)

func AuthenticatedUser(function CustomHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err, val := IsTokenValid(w, r)
		fmt.Println("after token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(err.Error())
			return
		}
		fmt.Println("before get r context")

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user-id", val)
		fmt.Println("after ser value to context")
		fmt.Println("after se with value to context")
		fmt.Println(ctx.Value("user-id"))
		function(w, r.WithContext(ctx))
	})
}

func IsUserAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	err, _ := IsTokenValid(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w)
	return

}
