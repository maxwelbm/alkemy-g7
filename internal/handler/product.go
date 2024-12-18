package handler

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
)



type ProductHandler struct {
	ProductService interfaces.IProductService
}

func NewProductHandler(ps interfaces.IProductService) *ProductHandler {
	return &ProductHandler{ProductService: ps}
}

func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
    products, err := ph.ProductService.GetAllProducts()
    var responseBody map[string]interface{}

    if err != nil {
        responseBody = map[string]interface{}{
            "error": err.Error(),
        }
        response.JSON(w, http.StatusNotFound, responseBody)
        return
    }

    responseBody = map[string]interface{}{
        "data": products,
    }

    response.JSON(w, http.StatusOK, responseBody)
}