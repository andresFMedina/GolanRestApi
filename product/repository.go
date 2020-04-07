package products

import (
	"database/sql"
)

//Repository interface
type Repository interface {
	GetProductByID(productID int) (*Product, error)
}

type repository struct {
	db *sql.DB
}

//NewRepository function
func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

//GetProductById method
func (repo *repository) GetProductByID(productID int) (*Product, error) {
	const query = `SELECT id,product_code,product_name,COALESCE(description, ''),
				 standard_cost, list_price,
				 category
				 FROM products
				 WHERE id=?`

	row := repo.db.QueryRow(query, productID)
	product := &Product{}

	err := row.Scan(&product.ID, &product.ProductCode, &product.ProductName, &product.Description,
		&product.StandardCost, &product.ListPrice, &product.Category)

	if err != nil {
		panic(err)
	}

	return product, err

}
