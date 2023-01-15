package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)
var DbSingleton *sql.DB


const username string = "manuel"
const password string = "manuel"
const host string = "localhost"
const port int = 3306
const DbName string = "golang"

const userSchema string =

	`CREATE TABLE IF NOT EXISTS users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(30) NOT NULL,
	last_name VARCHAR(30),
	password VARCHAR(64) NOT NULL,
	email VARCHAR(40) NOT NULL UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP )`



// Init Posts Table schema
const postSchema string =

	`CREATE TABLE IF NOT EXISTS posts (
	id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
	user_id INT NOT NULL,
	title VARCHAR(50) NOT NULL,
	content VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
  CONSTRAINT fk_post_user
    FOREIGN KEY (user_id) REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE RESTRICT
 	);`

func CreateConnection (){

	connection, err := sql.Open("sqlite3", "forum.Db")
	if err != nil {
		panic(err)
	} else {

		DbSingleton = connection
		fmt.Println("Connection successfully to database.")
	}
}
func Ping ()  {
	if err := DbSingleton.Ping();  err != nil{
		panic(err)
	}
}
func generateUrl () string{
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port,DbName)
}

func CloseConnection()  {
	err := DbSingleton.Close()
	if err != nil {
		fmt.Println("Problem closing the database")
	}
	fmt.Println("Database disconnected..")
}

func CreateTables ()  {
	createTable("users", userSchema)
	createTable("posts", postSchema)

}

func createTable(tablename, schema string) {
	if !existsTable(tablename) {
		_, err := Execute(schema)
		if err != nil {
			fmt.Printf("Error on create table. %s",err)

		}
	}
}

func existsTable(tableName string) bool {
	sqlSentence := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s';", tableName)
	rows, _ := Query(sqlSentence)
	return rows.Next()
}

func Execute(query string, args ...interface{}) (sql.Result, error) {
	result , err := DbSingleton.Exec(query, args...)
	if err != nil{
		log.Printf("Error executing <Execute>. %s",err)

	}
	return result, err
}

func Query (query string, args ...interface{}) (*sql.Rows, error)  {
	rows, err := DbSingleton.Query(query, args...)
	if err != nil{
		log.Printf("Error executing <Query>. %s",err)
		return nil, err
	}
	defer rows.Close()
	return rows, err
}

