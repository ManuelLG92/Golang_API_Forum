package user_domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id string `sql:"type:VARCHAR(255)" json:"id" gorm:"primaryKey"`
	Name string `sql:"type:VARCHAR(100)" json:"name"`
	LastName string `sql:"type:VARCHAR(100)" json:"last_name" gorm:"column:last_name"`
	Email string `sql:"type:VARCHAR(150);not null;unique" json:"email"`
	Password string `sql:"type:VARCHAR(255)" json:"password"`
}