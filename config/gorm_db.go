package config

import (
	"database/sql"
	"golang.com/forum/user/domain"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"log"
)
var DbGorm *gorm.DB
func Connection()  *gorm.DB{
	dsn := "manuel:manuel@tcp(127.0.0.1:3306)/golang_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	dbFunction, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DbGorm = dbFunction
	return DbGorm
}

func CreateGormDatabase()  {
	err := Connection().AutoMigrate(&user_domain.User{})
	if err != nil {
		return 
	}
}

func CloseGormConnection()  {
	sqlDB, err := DbGorm.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
}
