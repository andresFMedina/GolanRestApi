package main

import (
	"net/http"

	"github.com/GolanRestApi/database"
	products "github.com/GolanRestApi/product"
	"github.com/go-chi/chi"
)

func main() {
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	var productRepository = products.NewRepository(databaseConnection)
	var productService products.Service
	productService = products.NewService(productRepository)

	r := chi.NewRouter()
	r.Mount("/products", products.MakeHTTPHandler(productService))
	http.ListenAndServe(":3000", r)
}
