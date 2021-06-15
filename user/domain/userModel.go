package user_domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id int8 `sql:"type:VARCHAR(255)" json:"id"`
	Name string `sql:"type:VARCHAR(100)" json:"name"`
	Surname string `sql:"type:VARCHAR(100)" json:"surname"`
	Email string `sql:"type:VARCHAR(150);not null;unique" json:"email"`
	Password string `sql:"type:VARCHAR(255)" json:"password"`

}
