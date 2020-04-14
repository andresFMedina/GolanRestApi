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

//ProductList model
type ProductList struct {
	Data         []*Product `json:"data"`
	TotalRecords int        `json:"totalRecords"`
}

//ProductTop Model
type ProductTop struct {
	ID          int     `json:"id"`
	ProductName string  `json:"productName"`
	Sold        float64 `json:"sold"`
}

//ProductTopResponse Model
type ProductTopResponse struct {
	Data      []*ProductTop `json:"data"`
	TotalSold float64       `json:"totalSold"`
}
