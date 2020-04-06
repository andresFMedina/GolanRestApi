package main

import (
	"GolanRestApi/database"
	"fmt"
)

func main() {
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	fmt.Println(databaseConnection)
}
