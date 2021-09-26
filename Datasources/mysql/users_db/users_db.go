package users_db

import (
	"database/sql"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Below implementation can be used for Data Security of Database Credentials by storing the same in Environment Variables

/*const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host  ="mysql_users_host"
	mysql_users_schema ="mysql_users_schema"
)

var (
	Client *sql.DB
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host = os.Getenv(mysql_users_host)
	schema = os.Getenv(mysql_users_schema)

)*/

//Create database object to handle DB functions
var (
	Client *sql.DB
)

//To establish connection to the MySQL connection
func init(){
	
	var err error
	Client,err =sql.Open("mysql","root:rahima@tcp(127.0.0.1:3306)/users_db")
	if err!=nil{
		panic(err)
	}

	if err = Client.Ping();err!=nil{
		panic(err)
	}
	log.Println(("database Sucessfully Configured"))

}