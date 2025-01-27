package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	responses "github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customError"
)

type ProductRecHandler struct {
	ProductRecServ interfaces.IProductRecService
}

func NewProductRecHandler(prs interfaces.IProductRecService) *ProductRecHandler {
	return &ProductRecHandler{ProductRecServ: prs}
}

// CreateProductRecServ creates a new product record.
// @Summary Create a new product record
// @Description This endpoint allows for creating a new product record.
// @Tags ProductRecord
// @Produce json
// @Param productRecord body model.ProductRecords true "Product Record Information"
// @Success 201 {object} model.ProductRecordResponseSwagger{data=model.ProductRecords}
// @Failure 422 {object} model.ErrorResponseSwagger "Invalid JSON format"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to create product record"
// @Router /product-records [post]
func (prh *ProductRecHandler) CreateProductRecServ(w http.ResponseWriter, r *http.Request) {
	var productRecBody model.ProductRecords

	if err := json.NewDecoder(r.Body).Decode(&productRecBody); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("json mal formatado ou invalido", nil))
		return
	}

	product, err := prh.ProductRecServ.CreateProductRecords(productRecBody)
	if err != nil {
		if err, ok := err.(*customError.GenericError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody(customError.ErrUnknow.Error(), nil))
		return
	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", product))
}

// GetProductRecReport retrieves a product record report.
// @Summary Retrieve a product record report
// @Description This endpoint retrieves the product record report based on the provided product ID. If no ID is provided, it returns all records.
// @Tags ProductRecord
// @Produce json
// @Param id query int false "Product ID"
// @Success 200 {object} model.ProductRecordResponseSwagger{data=[]model.ProductRecords}
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid Parameter"
// @Failure 500 {object} model.ErrorResponseSwagger "Internal Server Error"
// @Router /product-records/report [get]
func (prh *ProductRecHandler) GetProductRecReport(w http.ResponseWriter, r *http.Request) {
	idProductStr := r.URL.Query().Get("id")
	idProduct := 0

	if idProductStr != "" {
		var err error

		idProduct, err = strconv.Atoi(idProductStr)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("Invalid Parameter", nil))
			return
		}
	}

	product, err := prh.ProductRecServ.GetProductRecordReport(idProduct)
	if err != nil {
		if appErr, ok := err.(*customError.GenericError); ok {
			response.JSON(w, appErr.Code, responses.CreateResponseBody(appErr.Error(), nil))
			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Internal Server Error", nil))
		return
	}

	if len(product) == 0 {
		response.JSON(w, http.StatusOK, responses.CreateResponseBody("empty list", nil))
		return
	}

	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", product))
}