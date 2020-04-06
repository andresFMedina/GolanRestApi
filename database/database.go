package database

import (
	"database/sql"
	"fmt"

	//Importando el modulo para la implementación de la cadena del sql
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	connectionString := "root:@tcp(localhost:3306)/northwind"
	fmt.Println(connectionString)

	databaseConnection, err := sql.Open("mysql", connectionString)	

	if err != nil {
		panic(err.Error()) //Error Handling
	}

	return databaseConnection
}
