package handlers

import (
	"sync"
	"time"
)

const (
	cookieName = "Go_session"
)


type UserAuthenticated struct {
	Id int `json:"id"`
	Name string `json:"name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
}

var (
	cookieLoginExpires = time.Now().Add(365 * 24 * time.Hour)
	syncSession = &sync.Mutex{}
)


// func SetSession(w http.ResponseWriter, r *http.Request, email string)  {
// 	syncSession.Lock()
// 	defer syncSession.Unlock()
// 	if getValueCookie(r) == "" {
// 		uuidCookie := uuid.NewV4()
// 		cookie := &http.Cookie{
// 			Name:    cookieName,
// 			Value:   uuidCookie.String(),
// 			Expires: cookieLoginExpires,
// 			Path:    "/",
// 		}
// 		http.SetCookie(w, cookie)
// 		userFound := GetUserByEmail(email)
// 		sqlInsertSession := "INSERT sessions SET uuid=?, user_id=?, name=?, last_name=?, email=?"
// 		_ , err := config.Execute(sqlInsertSession,uuidCookie.String(),
// 			userFound.Id,userFound.Name,userFound.LastName,email)
// 		if err != nil {
// 			//fmt.Println("Error inserting cookie: ", err)
// 			return
// 		} else {
// 			//fmt.Println("Cookies inserted: ", result)
// 		}
// 	} else {
// 		if lookSessionErr := getSessionData( getValueCookie(r)); lookSessionErr != nil {
// 			fmt.Println(lookSessionErr)
// 			return
// 		}
// 	}

// }

// func getValueCookie (r *http.Request) string {
// 	cookie, err := r.Cookie(cookieName)
// 	//fmt.Println(err)
// 	if err == nil {
// 		return cookie.Value
// 	}
// 	return ""
// }

// func getSessionData(cookieUuid string) *UserAuthenticated {
// 	authUser := UserAuthenticated{}
// 	sqlSearchSession := "SELECT user_id, name, last_name, email FROM sessions WHERE uuid=?"
// 	rows, err := config.Query(sqlSearchSession, cookieUuid)
// 	if err != nil {
// 		return nil
// 	}
// 	for rows.Next() {
// 		_=rows.Scan(&authUser.Id,&authUser.Name, &authUser.LastName, &authUser.Email)
// 		//return errors.New("Session already on database")
// 	}
// 	//rows.Close()
// 	return &authUser
// }


// func IsAuthenticated(r *http.Request) *UserAuthenticated {

// 	cookieClient, err := r.Cookie(cookieName)

// 	if err != nil {
// 		fmt.Println("Cookie not found ", err)
// 		return nil
// 	}

// 	userData := getSessionData(cookieClient.Value)
// 		if userData == nil {
// 			fmt.Println("Hubo un problema autenticando al usuario")
// 			return nil
// 		}
// 	return userData
// }

// func DeleteSession(w http.ResponseWriter, r *http.Request) error {
// 	err := deleteSessionFromDB(r)
// 	if err != nil {
// 		return err
// 	}
// 	//fmt.Println("Error from deleteSessionFromDB line 89: ", err)
// 	cookie := &http.Cookie{
// 		Name: cookieName,
// 		Value: "",
// 		Path: "/",
// 		MaxAge: -1,
// 	}
// 	http.SetCookie(w, cookie)
// 	return nil
// }

// func deleteSessionFromDB(r *http.Request ) error {
// 	cookie := getValueCookie(r)
// 	sql := "DELETE FROM sessions WHERE uuid=?"
// 	_, err := config.Execute(sql, cookie)
// 	if err != nil {
// 		fmt.Println("Error in line 106 execure delete cookie: ", err)
// 		return errors.New("Error deleting cookie")
// 	}
// 	return nil
// }




