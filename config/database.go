package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"

)
var db *sql.DB

const username string = "root"
const password string = "manuel92"
const host string = "localhost"
const port int = 3307
const dbName string = "go_react"

const userSchema string =

	`CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(30) NOT NULL,
	last_name VARCHAR(30),
	password VARCHAR(64) NOT NULL,
	email VARCHAR(40) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP )`



// Init Posts Table schema
const postSchema string =

	`CREATE TABLE posts (
	id INT(6) UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
	user_id INT(6) UNSIGNED NOT NULL,
	title VARCHAR(50) NOT NULL,
	content VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
  CONSTRAINT fk_post_user
    FOREIGN KEY (user_id) REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE RESTRICT
 	);`

// Init Sessions Table schema
const sessionSchema string =

	`CREATE TABLE sessions (
	uuid VARCHAR(100) NOT NULL PRIMARY KEY,
	email VARCHAR(40) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
 	)`


// End Sessions Table schema

//-- Init Database Connection
func CreateConnection (){
	//dbData := "root:manuel92@tcp(127.0.0.1:3307)/go"
	connection, err := sql.Open("mysql",generateUrl())
	if err != nil {
		//log.Fatal(err)
		panic(err)
	} else {
		db = connection
		fmt.Println("Connection successfully to database.")
	}
}
//-- End Database Connection

//-- Init Ping Database
func Ping ()  {
	if err := db.Ping();  err != nil{
		panic(err)
	}
}
//-- end Ping Database

//-- Init generateUIRL for Database Connection
func generateUrl () string{
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port,dbName)
}
//-- End generateUIRL for Database Connection

//-- Init Close Database Connection

func CloseConnection()  {
	err := db.Close()
	if err != nil {
		fmt.Println("Problem closing the database")
	}
	fmt.Println("Database disconnected..")
}
//-- End Close Database Connection

//-- Init  Create DB tables
func CreateTables ()  {
	createTable("users", userSchema)
	createTable("posts", postSchema)
	createTable("sessions", sessionSchema)

}
//-- End  Create DB tables

//-- Init  Create DB table
func createTable(tablename, schema string) {
	if !existsTable(tablename) {
		_, err := Execute(schema)
		if err != nil {
			fmt.Println(err)

		}
	}
}
//-- End  Create DB table

//-- Init  Check if exists DB table name
func existsTable(tableName string) bool {
	sqlSentence := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, _ := Query(sqlSentence)
	return rows.Next()
}
//-- End  Check if exists DB table name

func Execute(query string, args ...interface{}) (sql.Result, error) { //db.Exec back a result set and an error
	result , err := db.Exec(query, args...)
	if err != nil{
		log.Println(err)
	}
	return result, err
}

func Query (query string, args ...interface{}) (*sql.Rows, error)  {
	rows, err := db.Query(query, args...)
	if err != nil{
		log.Println(err)
		//return nil, err
	}
	//defer rows.Close()
	return rows, err
}

/*
func CreatePostTable()  {
	createTable(postTableName, postSchema)
}*/

//End Posts Table

