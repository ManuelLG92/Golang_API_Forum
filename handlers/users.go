package handlers

import (
	"encoding/json"
	"fmt"
	"golang.com/forum/models"
	"net/http"
	"strings"
	validator "github.com/go-playground/validator/v10"
	"log"
)


type User struct {
	Id int8 `json:"id" `
	Name string `json:"name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
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

func newUser (name, lastName, password, email string) *User {
	user := &User{Name: name, LastName: lastName, Password: password, Email: email} // Creamos un objeto user o instancia en Java referenciando con &
	//user.SetPassword(password)
	if  user.Valid() != nil {
		return nil
	}
	return user
}

// Init
func Index(w http.ResponseWriter, _ *http.Request)  {
	_, err := fmt.Fprintf(w, "You are in golang app!")
	if err != nil {
		return
	}
}
// End

// Init DONE
func SingUp (w http.ResponseWriter, r *http.Request)  {
	// w.Header().Set("Content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		user := User{}
		//decoder := json.NewDecoder(r.Body)
		fmt.Println(r.Body)
		//err := decoder.Decode(&user)
		validate := validator.New()
		err := validate.Struct(user)
		if err != nil {
			// log out this error
			log.Printf("Error validaing user %v",err)
			// return a bad request and a helpful error message
			// if you wished, you could concat the validation error into this
			// message to help point your consumer in the right direction.
			http.Error(w, "failed to validate struct", 400)
			//models.SendCustom(w, "body invalido", 422)
			return
		}
		//fmt.Println(err)
		if err != nil {
			fmt.Println("no se reconoce lo enviado por el usuario")
			models.SendCustom(w, "body invalido", 422)
		} else {
			fmt.Println("REconocido JSON enviado por el cliente")
			fmt.Println(user)
			userCreated , err := CreateUser(strings.TrimSpace(user.Name), strings.TrimSpace(user.LastName),
				strings.TrimSpace(user.Password), strings.TrimSpace(user.Email))
			if err != nil {
				fmt.Println("Error al tratar de crear al usuario: ", err)
				models.SendNoContent(w)
			} else {
				fmt.Println("usuario creado")
				err := json.NewEncoder(w).Encode(userCreated)
				if err != nil {
					return
				}

			}
		}
}
// End

// Init CreateUser
func CreateUser(name, lastName, password, email string) (*User, error) {
	user := newUser(name, lastName, password, email)
	err := user.Valid() //Data validation
	fmt.Println("El error es: ",err)
	checkMail := existEmail(email)
	if checkMail != nil {
		return nil, checkMail
	}
	if err != nil {
		return nil, err //Return error
	} else {
		err := user.SetPassword(password)
		if err != nil {
			return nil, err
		} //Hash password
		errInsertUSer := user.insertUser() //Insert user in BBDD
		if errInsertUSer != nil {
			fmt.Println("Ha ocurrido: ",errInsertUSer)
			return nil, errInsertUSer
		} else {
			return user, nil //Return user Created
		}

	}
}
//End CreateUser


// Init SignIn DONE
func SingIn (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	//err, user := decodeUser(r)
	/**/user := User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	fmt.Println(user)
	if err != nil {
		fmt.Println("no se reconoce lo enviado por el usuario")
		//w.WriteHeader(http.StatusNotAcceptable)
		//json.NewEncoder(w).Encode("Json format unrecognized ")
		//return

	} else {
		fmt.Println("REconocido JSON enviado por el cliente")
		//fmt.Println(user)
		emailUser := user.Email
		user.Email = strings.TrimSpace(user.Email)
		user.Password = strings.TrimSpace(user.Password)
		response, user := login(user.Email, user.Password)
		if response == false {
			fmt.Println("Error al tratar de validar el login: ", err)
			models.SendNotAuth(w)
		} else {
			fmt.Println("usuario autenticado")
			SetSession(w, r, emailUser)
			models.SendData(w, user)
			//json.NewEncoder(w).Encode(userCreated)
			//models.SendData(w,userCreated)
		}
	}


}
//End SignIn

// Init update user

func UpdateUsers (w http.ResponseWriter, _ *http.Request)  {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
//End UpdateUser

// Init
func GetUserbyId(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	userId := getIdByUrl(r)
	if userId != 0 {

	}

}

// End
// Init Valid Data
func (user *User) Valid() error  {

	if  validEmail(user.Email) != nil {
		return models.ErrorEmail
		//fmt.Println("Invalid Email")
	}
	if validData(user.Name, user.LastName, user.Password)!= nil {
		return models.ErrorNotValidData
		//fmt.Println("Invalid names")
		//return ErrorNotValidUser
	}
	return nil
}
//End Valid Data


//Init/
// End

// Init SignIn DONE
func LogOut (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	//c, _ := r.Cookie("Go_session")
	//fmt.Println(c.Value)
	err := DeleteSession(w, r)
	if err != nil {
		fmt.Println("Error in line 161 Logout main: ",err)
		models.SendNotAuth(w)
	} else {
		models.SendData(w, "Session Deleted")
	}
/*
	user := User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	fmt.Println(user)
	if err != nil {
		fmt.Println("no se reconoce lo enviado por el usuario")
	} else {
		fmt.Println("REconocido JSON enviado por el cliente")
		//fmt.Println(user)
		user.Email = strings.TrimSpace(user.Email)
		user.Password = strings.TrimSpace(user.Password)
		response, user := login(user.Email, user.Password)
		if response == false {
			fmt.Println("Error al tratar de validar el login: ", err)
			models.SendNotAuth(w)
		} else {
			fmt.Println("usuario autenticado")
			config.SetSession(w)
			models.SendData(w, user)
			//json.NewEncoder(w).Encode(userCreated)
			//models.SendData(w,userCreated)
		}*/
	}


//End SignIn


// Init
func GetUsers (w http.ResponseWriter,_ *http.Request)  {
	_, _ =fmt.Fprintf(w, "Welcome to GetUsers route!")
}
// End
