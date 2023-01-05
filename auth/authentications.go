package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CustomHandler func(w http.ResponseWriter, r *http.Request)

type Middleware struct {
	next http.Handler
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// run your handler code here
	// write error into w and return if you need to interrupt request execution

	// call next handler
	m.next.ServeHTTP(w, r)
}

func AuthenticatedUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		timeStart := ctx.Value("start-at").(time.Time)
		fmt.Println("request started at", timeStart.String())
		err, val := IsTokenValid(w, r)
		fmt.Println("after token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx = context.WithValue(ctx, "user-id", val)

		next(w, r.WithContext(ctx))
	})
}

func StartRequest(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "start-at", time.Now())
		defer fmt.Println(time.Now().String() + "after")
		next(w, r.WithContext(ctx))

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
