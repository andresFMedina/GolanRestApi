package main

import (
	"GolanRestApi/database"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

var databaseConnection *sql.DB

// Product type
type Product struct {
	ID          int    `json:"id"`
	ProductCode string `json:"product_code"`
	Description string `json:"description"`
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	databaseConnection = database.InitDB()
	defer databaseConnection.Close()

	r := chi.NewRouter()
	r.Get("/products", AllProducts)
	r.Post("/products", CreateProducts)
	r.Put("/products/{id}", UpdateProduct)
	r.Delete("/products/{id}", DeleteProduct)
	http.ListenAndServe(":3000", r)
}

//AllProducts retorna todos los productos
func AllProducts(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id, product_code, COALESCE(description, '')
				 FROM products`
	results, err := databaseConnection.Query(sql)
	catch(err)
	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.ProductCode, &product.Description)

		catch(err)
		products = append(products, product)
	}
	respondwithJSON(w, http.StatusOK, products)
}

//CreateProducts post products
func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var product Product
	json.NewDecoder(r.Body).Decode(&product)

	query, err := databaseConnection.Prepare("Insert products SET product_code=?, description=?")
	catch(err)

	_, er := query.Exec(product.ProductCode, product.Description)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "success"})
}

//UpdateProduct actualiza el producto
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&product)

	query, err := databaseConnection.Prepare("Update products set product_code=?, description=? where id=?")
	catch(err)
	_, er := query.Exec(product.ProductCode, product.Description, id)
	catch(er)

	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "success"})
}

//DeleteProduct Borra el producto
func DeleteProduct(w http.ResponseWriter, r *http.Request){	
	id := chi.URLParam(r, "id")	

	query, err := databaseConnection.Prepare("delete from products where id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)

	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "deleted"})

}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
