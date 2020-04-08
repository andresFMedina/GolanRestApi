package products

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

//MakeHTTPHandler method
func MakeHTTPHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getProductByIDHandler := kithttp.NewServer(makeGetProductByIDEndPoint(s), getProductByIDRequestDecoder, kithttp.EncodeJSONResponse)

	r.Method(http.MethodGet, "/{id}", getProductByIDHandler)

	getProductsHandler := kithttp.NewServer(makeGetProductsEndPoint(s),
		getProductsRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProductsHandler)

	return r
}

func getProductByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIDRequest{
		ProductID: productID,
	}, nil
}

func getProductsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	return request, nil
}
