package handler

import (
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"net/http"
	"strconv"
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

func (ph *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	var responseBody map[string]interface{}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseBody = map[string]interface{}{
			"error": "id invalido",
		}
		response.JSON(w, http.StatusBadRequest, responseBody)
		return
	}

	product, err := ph.ProductService.GetProductById(id)
	if err != nil {
		responseBody = map[string]interface{}{
			"error": err.Error(),
		}
		response.JSON(w, http.StatusNotFound, responseBody)
		return
	}

	responseBody = map[string]interface{}{
		"data": product,
	}

	response.JSON(w, http.StatusOK, responseBody)
}

func (ph *ProductHandler) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	var responseBody map[string]interface{}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		responseBody = map[string]interface{}{
			"error": "id invalido",
		}
		response.JSON(w, http.StatusBadRequest, responseBody)
		return
	}

	err = ph.ProductService.DeleteProduct(id)
	if err != nil {
		responseBody = map[string]interface{}{
			"error": err.Error(),
		}
		response.JSON(w, http.StatusNotFound, responseBody)
		return
	}

	responseBody = map[string]interface{}{
		"data": "Produto Deletado",
	}

	response.JSON(w, http.StatusNoContent, responseBody)
}
