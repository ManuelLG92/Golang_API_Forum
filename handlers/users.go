package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.com/forum/auth"
	"golang.com/forum/helpers"
	"golang.com/forum/models"
	"golang.com/forum/routes"
)

type User struct {
	Id       string `json:"id" `
	Name     string `json:"name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UserRoutes := []routes.Routes{}
var signUp routes.Routes = routes.Routes{Path: "/sign-up/", Name: "register", Methods: []string{"POST", "OPTIONS"}, Handler: SingUp, NeedsAuth: false}
var signIn routes.Routes = routes.Routes{Path: "/login/", Name: "login", Methods: []string{"POST", "OPTIONS"}, Handler: SingIn, NeedsAuth: false}

func GetRoutes() *[]routes.Routes {
	var signUp routes.Routes = routes.Routes{Path: "/sign-up/", Name: "register", Methods: []string{"POST", "OPTIONS"}, Handler: SingUp, NeedsAuth: false}
	var signIn routes.Routes = routes.Routes{Path: "/login/", Name: "login", Methods: []string{"POST", "OPTIONS"}, Handler: SingIn, NeedsAuth: false}
	var index routes.Routes = routes.Routes{Path: "/", Name: "login", Methods: []string{"GET"}, Handler: Index, NeedsAuth: true}

	var UserRoutes []routes.Routes = []routes.Routes{}
	UserRoutes = append(UserRoutes, signUp)
	UserRoutes = append(UserRoutes, signIn)
	UserRoutes = append(UserRoutes, index)

	return &UserRoutes

}

/*
	{
		"name":"test",
		"last_name": "test",
		"email": "tes87@test.test",
		"password": "test1234"

}
*/
type Users []User

func (user *User) IsValidOnCreation() bool {
	if user.Email == "" || user.Name == "" || user.LastName == "" || user.Password == "" {
		return false
	}
	return true
}
func newUser(name, lastName, password, email string) *User {
	user := &User{Id: uuid.New().String(), Name: name, LastName: lastName, Password: password, Email: email} // Creamos un objeto user o instancia en Java referenciando con &
	//user.SetPassword(password)
	if user.Valid() != nil {
		return nil
	}
	return user
}

// Init
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("before ctx.")
	userCtx := r.Context().Value("user-id").(*string)
	fmt.Printf("user ctx. %v", *userCtx)
	_, err := fmt.Fprintf(w, "You are in golang app!")
	if err != nil {
		return
	}
}

// End

// Init DONE
func SingUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var user *User

	// decoder := json.NewDecoder(r.Body)
	// fmt.Println(r.Body)
	// errDecoding := decoder.Decode(&user)

	user, err := helpers.DecodeBody[User](r.Body, "missing fields on user")
	if err != nil {
		log.Printf("Error parsing user %v", err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	if err = validator.New().Struct(user); err != nil {
		log.Printf("Error validaing user %v", err)
		http.Error(w, err.Error(), 422)
		return
	}
	userFetched, err := GetUserByEmail(user.Email)
	if userFetched != nil {
		fmt.Printf("user with email %v already registered", user.Email)
		models.SendCustom(w, fmt.Sprintf("user with email %v aleady registered", user.Email), 400)
		return
	}
	userCreated, err := CreateUser(strings.TrimSpace(user.Name), strings.TrimSpace(user.LastName),
		strings.TrimSpace(user.Password), strings.TrimSpace(user.Email))
	if err != nil {
		fmt.Println("Error al tratar de crear al usuario: ", err)
		models.SendNoContent(w)
		return
	}
	models.SendCreated(w, userCreated)
	return
}

func CreateUser(name, lastName, password, email string) (*User, error) {
	user := newUser(name, lastName, password, email)
	err := user.Valid() //Data validation
	if err != nil {
		fmt.Printf("error on validation: %v", err.Error())
		return nil, models.InterServerError
	}
	exists, err := existEmail(email)
	if err != nil {
		fmt.Printf("error: %v", err.Error())
		return nil, models.InterServerError
	}
	if *exists == false {
		fmt.Printf("error: %v", models.ErrorUserRegistred.Error())
		return nil, models.ErrorUserRegistred
	}
	err = user.SetPassword(password)
	if err != nil {
		fmt.Println("Unable to has the password")
		return nil, err
	}
	errInsertUSer := user.insertUser()
	if errInsertUSer != nil {
		fmt.Println("Unexpected error trying to save user. ", errInsertUSer.Error())
		return nil, errInsertUSer
	}
	return user, nil //Return user Created

}

func SingIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var user *User

	user, err := helpers.DecodeBody[User](r.Body, "missing fields on user")
	if err != nil {
		log.Printf("Error parsing user %v", err.Error())
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println("REconocido JSON enviado por el cliente")
	//fmt.Println(user)
	//emailUser := user.Email
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
	user, err = login(user.Email, user.Password)
	if err != nil {
		fmt.Println("Password and/or email invalid", err.Error())
		models.SendNotAuth(w)
		return
	}

	err, token := auth.GenerateJwt(auth.JwtCustomClaims{Id: user.Id, Name: user.Name, Email: user.Email, Password: user.Password, IP: r.RemoteAddr})
	if err != nil {
		fmt.Printf("Error creating token. %v", err.Error())
		models.SendInternalServerError(w)
		return
	}
	fmt.Println("authenticated user")
	w.Header().Set("x-access-token", *token)
	w.WriteHeader(http.StatusOK)
}

func UpdateUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

//End UpdateUser

// Init
func GetUserbyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	userId := getIdByUrl(r)
	if userId != "" {
		w.WriteHeader(http.StatusOK)
	}

}

// End
// Init Valid Data
func (user *User) Valid() error {

	if validEmail(user.Email) != nil {
		return models.ErrorEmail
		//fmt.Println("Invalid Email")
	}
	if validData(user.Name, user.LastName, user.Password) != nil {
		return models.ErrorNotValidData
		//fmt.Println("Invalid names")
		//return ErrorNotValidUser
	}
	return nil
}

// Init
func GetUsers(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome to GetUsers route!")
}

// End
