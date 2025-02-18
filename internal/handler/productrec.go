package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	responses "github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger" // Importando o pacote de logger
)

type ProductRecHandler struct {
	ProductRecServ interfaces.IProductRecService
	log            logger.Logger
}

func NewProductRecHandler(prs interfaces.IProductRecService, logger logger.Logger) *ProductRecHandler {
	return &ProductRecHandler{
		ProductRecServ: prs,
		log:           logger, // Inicializando o logger
	}
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
	prh.log.Log("ProductRecHandler", "INFO", "CreateProductRecServ function initializing")

	var productRecBody model.ProductRecords

	if err := json.NewDecoder(r.Body).Decode(&productRecBody); err != nil {
		prh.log.Log("ProductRecHandler", "ERROR", "Invalid JSON provided: "+err.Error())
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("json mal formatado ou invalido", nil))
		return
	}

	product, err := prh.ProductRecServ.CreateProductRecords(productRecBody)
	if err != nil {
		if appErr, ok := err.(*customerror.GenericError); ok {
			prh.log.Log("ProductRecHandler", "ERROR", fmt.Sprintf("Error creating product record: %s", appErr.Error()))
			response.JSON(w, appErr.Code, responses.CreateResponseBody(appErr.Error(), nil))
			return
		}

		prh.log.Log("ProductRecHandler", "ERROR", "Unable to create product record: "+customerror.ErrUnknow.Error())
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody(customerror.ErrUnknow.Error(), nil))
		return
	}

	prh.log.Log("ProductRecHandler", "INFO", "Product record created successfully")
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
	prh.log.Log("ProductRecHandler", "INFO", "GetProductRecReport function initializing")

	idProductStr := r.URL.Query().Get("id")
	idProduct := 0

	if idProductStr != "" {
		var err error
		idProduct, err = strconv.Atoi(idProductStr)
		if err != nil {
			prh.log.Log("ProductRecHandler", "ERROR", "Invalid parameter: Product ID cannot be converted to int")
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("Invalid Parameter", nil))
			return
		}
	}

	product, err := prh.ProductRecServ.GetProductRecordReport(idProduct)
	if err != nil {
		if appErr, ok := err.(*customerror.GenericError); ok {
			prh.log.Log("ProductRecHandler", "ERROR", fmt.Sprintf("Error getting product record report: %s", appErr.Error()))
			response.JSON(w, appErr.Code, responses.CreateResponseBody(appErr.Error(), nil))
			return
		}

		prh.log.Log("ProductRecHandler", "ERROR", "Internal Server Error occurred")
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Internal Server Error", nil))
		return
	}

	if len(product) == 0 {
		prh.log.Log("ProductRecHandler", "INFO", "No records found for Product ID: "+strconv.Itoa(idProduct))
		response.JSON(w, http.StatusOK, responses.CreateResponseBody("empty list", nil))
		return
	}

	prh.log.Log("ProductRecHandler", "INFO", "Product record report retrieved successfully")
	response.JSON(w, http.StatusOK, responses.CreateResponseBody("", product))
}