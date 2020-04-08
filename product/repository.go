package products

import (
	"database/sql"
)

//Repository interface
type Repository interface {
	GetProductByID(productID int) (*Product, error)
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *getUpdateProductRequest) (int, error)
}

type repository struct {
	db *sql.DB
}

//NewRepository function
func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

//GetProducts Method
func (repo *repository) GetProducts(params *getProductsRequest) ([]*Product, error) {
	const query = `SELECT id,product_code,product_name,COALESCE(description, ''),
				   standard_cost, list_price,
				   category
				   FROM products
				   LIMIT ? OFFSET ?`
	results, err := repo.db.Query(query, params.Limit, params.Offset)
	if err != nil {
		panic(err)
	}

	var products []*Product
	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.ProductCode, &product.ProductName, &product.Description,
			&product.StandardCost, &product.ListPrice, &product.Category)

		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	return products, err
}

func (repo *repository) GetTotalProducts() (int, error) {
	const sql = "SELECT COUNT(*) FROM PRODUCTS"
	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)

	if err != nil {
		panic(err)
	}

	return total, nil
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

//InsertProduct method
func (repo *repository) InsertProduct(params *getAddProductRequest) (int64, error) {
	const query = `INSERT INTO products
					 (product_code,product_name,description,
					 standard_cost, list_price,
					 category)
					 VALUES(?,?,?,?,?,?)`

	result, err := repo.db.Exec(query, params.ProductCode, params.ProductName, params.Description,
		params.StandardCost, params.ListPrice, params.Category)

	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	return id, nil

}

//UpdateProduct method
func (repo *repository) UpdateProduct(params *getUpdateProductRequest) (int, error) {
	const query = `
		UPDATE products
		SET product_code = ?,
		product_name = ?,
		description = ?,
		standard_cost = ?,
		list_price = ?,
		category = ?
		WHERE id = ?	
	`

	_, err := repo.db.Exec(query, params.ProductCode, params.ProductName, params.Description,
		params.StandardCost, params.ListPrice, params.Category, params.ID)

	if err != nil {
		panic(err)
	}

	id := params.ID

	return id, nil
}
