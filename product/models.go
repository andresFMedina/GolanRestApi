package products

//Product model
type Product struct {
	ID           int     `json:"id"`
	ProductCode  string  `json:"productCode"`
	ProductName  string  `json:"productName"`
	Description  string  `json:"description"`
	StandardCost float64 `json:"standardCost"`
	ListPrice    float64 `json:"listPrice"`
	Category     string  `json:"category"`
}
