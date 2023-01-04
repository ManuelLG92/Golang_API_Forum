package config

import (
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"log"
)
var DbGorm *gorm.DB
func Connection()  *gorm.DB{
	//dsn := "manuel:manuel@tcp(127.0.0.1:3306)/golang_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//dbFunction, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	dbFunction, err := gorm.Open(sqlite.Open("gorm-new.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DbGorm = dbFunction
	return DbGorm
}

func CloseGormConnection()  {
	sqlDB, err := DbGorm.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
		log.Fatalln(err)
		}
	}(sqlDB)
}
