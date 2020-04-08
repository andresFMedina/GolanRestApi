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


func makeGetProductsEndPoint(s Service) endpoint.Endpoint{
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