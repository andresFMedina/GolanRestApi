package products

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int
}

type getProductsRequest struct {
	Limit  int
	Offset int
}

type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    float64
	StandardCost float64
	ProductCode  string
	ProductName  string
}

type getUpdateProductRequest struct {
	ID           int
	Category     string
	Description  string
	ListPrice    float64
	StandardCost float64
	ProductCode  string
	ProductName  string
}

type deleteProductRequest struct {
	ProductID int
}

type getBestSellers struct {
	
}

func makeGetProductByIDEndPoint(s Service) endpoint.Endpoint {
	getProductByIDEndPoint := func(context context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)
		product, err := s.GetProductByID(&req)
		if err != nil {
			panic(err)
		}
		return product, nil
	}

	return getProductByIDEndPoint
}

func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(context context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return getProductsEndPoint
}

func makeAddProductEndPoint(s Service) endpoint.Endpoint {
	addProductEndPoint := func(context context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		result, err := s.InsertProduct(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return addProductEndPoint
}

func makeUpdateProductEndPoint(s Service) endpoint.Endpoint {
	updateProductEndPoint := func(context context.Context, request interface{}) (interface{}, error) {
		req := request.(getUpdateProductRequest)
		result, err := s.UpdateProduct(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}
	return updateProductEndPoint
}

func makeDeleteProductEndPoint(s Service) endpoint.Endpoint {
	deleteProductEndPoint := func(context context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductRequest)
		result, err := s.DeleteProduct(&req)
		if err != nil {
			panic(err)
		}
		return result, nil
	}

	return deleteProductEndPoint
}

func makeGetBestSellersEndPoint(s Service) endpoint.Endpoint {
	getBestSellersEndPoint := func(context context.Context, request interface{}) (interface{}, error) {
		result, err := s.GetBestSellers()
		if err != nil {
			panic(err)
		}
		return result, nil
	}

	return getBestSellersEndPoint
}
