package userapplication

import (
	"encoding/json"
	db "golang.com/forum/config"
	userDomain "golang.com/forum/user/domain"
	"gorm.io/gorm"
	"net/http"
)

type handleUser struct {
	userDomain.User
}

func  CreateUser(user userDomain.User) *gorm.DB{
	result := db.DbGorm.Create(&user)
	return result

}

func getAllUsers(w http.ResponseWriter, r *http.Request)  {
	users := db.DbGorm.Find(&userDomain.User{})
	w.Header().Set("Content-Type", "application/json")

	_= json.NewEncoder(w).Encode(users)
}