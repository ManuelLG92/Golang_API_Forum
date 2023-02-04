package auth

import (
	"fmt"
	"net/http"
	"time"

	"forum/helpers"
	jwt "github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Id    string
	Name  string
	Email string
	IP    string
	jwt.RegisteredClaims
}

type Claims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("my_secret_key")

func GenerateJwt(data JwtCustomClaims) (error, *string) {

	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &JwtCustomClaims{
		data.Id,
		data.Name,
		data.Email,
		data.IP,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedStringToken, err := token.SignedString(jwtKey)
	if err != nil {
		return err, nil
	}
	return nil, &signedStringToken
}

func IsTokenValid(_ http.ResponseWriter, r *http.Request) (error, *string) {
	claims := &Claims{}
	invalidToken := helpers.InvalidToken
	var token = r.Header.Get("x-access-token")
	if token == "" || len(token) < 40 {
		fmt.Println("invalidToken length or not present")
		fmt.Println(invalidToken.Error())
		return invalidToken, nil
	}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		fmt.Println("error invalidToken", token)
		fmt.Println(invalidToken.Error())
		return invalidToken, nil
	}
	if !tkn.Valid {
		fmt.Println(invalidToken.Error())
		return invalidToken, nil
	}

	fmt.Printf("Claim id. %v", claims.Id)
	return nil, &claims.Id
}
