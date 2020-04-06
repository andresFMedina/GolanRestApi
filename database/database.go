package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	connectionString := "root:admin@tcp(localhost:3306)/northwind"
	fmt.Println(connectionString)

	databaseConnection, err := sql.Open("mysql", connectionString)	

	if err != nil {
		panic(err.Error()) //Error Handling
	}

	return databaseConnection
}
