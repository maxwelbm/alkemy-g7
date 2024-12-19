package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	responses "github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
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

	if err != nil {
		response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("Success", products))
}

func (ph *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("id invalido", nil))
		return
	}

	product, err := ph.ProductService.GetProductById(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("Success", product))
}

func (ph *ProductHandler) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {

		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("id invalido", nil))
		return
	}

	err = ph.ProductService.DeleteProduct(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusNoContent, responses.CreateResponseBody("produto deletado", nil))
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productBody model.Product

	if err := json.NewDecoder(r.Body).Decode(&productBody); err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("json mal formatado ou invalido", nil))
	}

	product, err := ph.ProductService.CreateProduct(productBody)
	if err != nil {
		response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("success", product))
}

func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
		return
	}
	var productBody model.Product

	if err := json.NewDecoder(r.Body).Decode(&productBody); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("json mal formatado ou invalido", nil))
	}

	product, err := ph.ProductService.UpdateProduct(id, productBody)
	if err != nil {
		response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("success", product))
}
