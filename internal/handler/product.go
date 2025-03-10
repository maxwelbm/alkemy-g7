package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	responses "github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type ProductHandler struct {
	ProductService interfaces.IProductService
	log            logger.Logger
}

func NewProductHandler(ps interfaces.IProductService, logger logger.Logger) *ProductHandler {
	return &ProductHandler{
		ProductService: ps,
		log:            logger,
	}
}

// GetAllProducts retrieves all products.
// @Summary Retrieve all products
// @Description Fetch all registered products from the database
// @Tags Product
// @Produce json
// @Success 200 {object} model.ProductResponseSwagger
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to list products"
// @Router /products [get]
func (ph *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	ph.log.Log("ProductHandler", "INFO", "GetAllProducts function initializing")

	products, err := ph.ProductService.GetAllProducts()

	if err != nil {
		ph.log.Log("ProductHandler", "ERROR", "Unable to retrieve products: "+err.Error())
		response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	ph.log.Log("ProductHandler", "INFO", "Successfully retrieved products")
	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", products))
}

// GetProductByID retrieves a product by its ID.
// @Summary Retrieve a product
// @Description This endpoint fetches the details of a specific product based on the provided product ID.
// @Tags Product
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.ProductResponseSwagger{data=model.Product}
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid ID"
// @Failure 404 {object} model.ErrorResponseSwagger "Product Not Found"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to search for product"
// @Router /products/{id} [get]
func (ph *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		ph.log.Log("ProductHandler", "ERROR", "Invalid ID provided in the request: "+chi.URLParam(r, "id"))
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id", nil))
		return
	}

	product, err := ph.ProductService.GetProductByID(id)

	if err != nil {
		ph.log.Log("ProductHandler", "ERROR", fmt.Sprintf("Product not found for ID: %d, error: %v", id, err))
		response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	ph.log.Log("ProductHandler", "INFO", fmt.Sprintf("Successfully retrieved product with ID: %d", id))
	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", product))
}

// DeleteProductByID deletes a product by its ID.
// @Summary Delete a product by ID
// @Description This endpoint allows for deleting a product based on the provided product ID.
// @Tags Product
// @Produce json
// @Param id path int true "Product ID"
// @Success 204 {object} nil "Product successfully deleted"
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid ID"
// @Failure 404 {object} model.ErrorResponseSwagger "Product not found"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to delete product"
// @Router /products/{id} [delete]
func (ph *ProductHandler) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		ph.log.Log("ProductHandler", "ERROR", "Invalid ID provided for deletion: "+chi.URLParam(r, "id"))
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id", nil))
		return
	}

	err = ph.ProductService.DeleteProduct(id)

	if err != nil {
		if appErr, ok := err.(*customerror.GenericError); ok {
			ph.log.Log("ProductHandler", "ERROR", fmt.Sprintf("Error deleting product with ID: %d, error: %s", id, appErr.Error()))
			response.JSON(w, appErr.Code, responses.CreateResponseBody(appErr.Error(), nil))
			return
		}

		ph.log.Log("ProductHandler", "ERROR", fmt.Sprintf("Unable to delete product with ID: %d, error: %v", id, err))
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Internal Server Error", nil))
		return
	}

	ph.log.Log("ProductHandler", "INFO", fmt.Sprintf("Product with ID: %d successfully deleted", id))
	response.JSON(w, http.StatusNoContent, responses.CreateResponseBody("product deleted", nil))
}

// CreateProduct creates a new product.
// @Summary Create a new product
// @Description This endpoint allows for creating a new product.
// @Tags Product
// @Produce json
// @Param product body model.Product true "Product information"
// @Success 201 {object} model.ProductResponseSwagger{data=model.Product}
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid input"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to create product"
// @Router /products [post]
func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productBody model.Product

	if err := json.NewDecoder(r.Body).Decode(&productBody); err != nil {
		ph.log.Log("ProductHandler", "ERROR", "Invalid JSON syntax: "+err.Error())
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("invalid json syntax", nil))
		return
	}

	product, err := ph.ProductService.CreateProduct(productBody)

	if err != nil {
		ph.log.Log("ProductHandler", "ERROR", "Unable to create product: "+err.Error())
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	ph.log.Log("ProductHandler", "INFO", "Product created successfully")
	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", product))
}

// UpdateProduct updates an existing product.
// @Summary Update an existing product
// @Description This endpoint allows for updating the details of a specific product identified by the provided ID.
// @Tags Product
// @Produce json
// @Param id path int true "Product ID"
// @Param product body model.Product true "Product information"
// @Success 200 {object} model.ProductResponseSwagger{data=model.Product} "Product successfully updated"
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid ID"
// @Failure 404 {object} model.ErrorResponseSwagger "Product not found"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to update product"
// @Router /products/{id} [patch]
func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		ph.log.Log("ProductHandler", "ERROR", "Invalid ID provided for update: "+chi.URLParam(r, "id"))
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid id", nil))
		return
	}

	var productBody model.Product

	if err := json.NewDecoder(r.Body).Decode(&productBody); err != nil {
		ph.log.Log("ProductHandler", "ERROR", "Invalid JSON syntax: "+err.Error())
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("invalid json syntax", nil))
		return
	}

	product, err := ph.ProductService.UpdateProduct(id, productBody)

	if err != nil {
		ph.log.Log("ProductHandler", "ERROR", fmt.Sprintf("Unable to update product with ID: %d, error: %v", id, err))
		response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	ph.log.Log("ProductHandler", "INFO", fmt.Sprintf("Product with ID: %d updated successfully", id))
	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", product))
}