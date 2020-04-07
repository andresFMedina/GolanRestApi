package products

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getProductByIDRequest struct {
	ProductID int
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
