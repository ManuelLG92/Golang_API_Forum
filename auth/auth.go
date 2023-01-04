package auth

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.com/forum/models"
)

type JwtCustomClaims struct {
	Id          string  
	Name        string 
	Email        string 
	Password        string 
	IP          string 
	jwt.StandardClaims
}


type Claims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("my_secret_key")
func GenerateJwt(data JwtCustomClaims) (error, *string) {

	claims := &JwtCustomClaims{
		data.Id,
		data.Name,
		data.Email,
		data.Password,
		data.IP,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
		},
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)


	signedStringToken, err := token.SignedString(jwtKey)
	if err != nil {
		return err, nil
	}
	return nil, &signedStringToken
}

func IsTokenValid(w http.ResponseWriter, r *http.Request) (error, interface{})  {
	claims := &Claims{}
	invalidToken := models.InvalidToken
	var token = r.Header.Get("x-access-token")
	fmt.Printf("Token %v", token)
	if token == "" || len(token) < 40 {
		fmt.Println("invalidToken length or not present")
		fmt.Println(invalidToken.Error())
		return invalidToken, nil
	}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		fmt.Println("error invalidToken")
	fmt.Println(invalidToken.Error())
		return invalidToken, nil
	}
	if  !tkn.Valid {
	   fmt.Println(invalidToken.Error())
		return invalidToken, nil
	}
	return nil, claims.Id
}
