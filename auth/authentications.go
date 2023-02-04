package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const ContextUserId = "user-id"
const ContextStartAt = "start-at"

func AuthenticatedUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		timeStart := ctx.Value(ContextStartAt).(time.Time)
		fmt.Println("request started at", timeStart.String())
		err, val := IsTokenValid(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx = context.WithValue(ctx, ContextUserId, val)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func StartRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("req body start request", r.Body)
		ctx := r.Context()
		ctx = context.WithValue(ctx, ContextStartAt, time.Now())
		defer fmt.Println(time.Now().String() + "after")
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

//func IsUserAuth(w http.ResponseWriter, r *http.Request) {
//	err, _ := IsTokenValid(w, r)
//	if err != nil {
//		w.WriteHeader(http.StatusUnauthorized)
//		_ = json.NewEncoder(w).Encode(err.Error())
//		return
//	}
//	w.WriteHeader(http.StatusAccepted)
//	_ = json.NewEncoder(w)
//	return
//
//}

func GetUserIdFromContext(ctx context.Context) *string {
	usr, err := ctx.Value("user-id").(*string)
	if !err {
		fmt.Println("error", err)
		return nil
	}
	fmt.Println("no error", usr, &usr)

	return usr
}
