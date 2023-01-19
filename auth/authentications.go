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

const ContextUserId = "user-id"
const ContextStartAt = "start-at"
func AuthenticatedUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		timeStart := ctx.Value(ContextStartAt).(time.Time)
		fmt.Println("request started at", timeStart.String())
		err, val := IsTokenValid(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx = context.WithValue(ctx,ContextUserId, val)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func StartRequest(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		fmt.Println("req body start request",r.Body)
		ctx := r.Context()
		ctx = context.WithValue(ctx,ContextStartAt, time.Now())
		defer fmt.Println(time.Now().String() + "after")
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func IsUserAuth(w http.ResponseWriter, r *http.Request) {
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

func GetUserIdFromContext(ctx context.Context) *string  {
	usr, err := ctx.Value("user-id").(*string)
	if !err {
		fmt.Println("error", err)
		return nil
	}
	fmt.Println("no error", usr, &usr)

	return usr
}
