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
	appErr "github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type ProductRecHandler struct {
	ProductRecServ interfaces.IProductRecRepo
}

func NewProductRecHandler(prs interfaces.IProductRecRepo) *ProductRecHandler {
	return &ProductRecHandler{ProductRecServ: prs}
}

func (prh *ProductRecHandler) CreateProductRecServ(w http.ResponseWriter, r *http.Request) {
	var productRecBody model.ProductRecords

	if err := json.NewDecoder(r.Body).Decode(&productRecBody); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("json mal formatado ou invalido", nil))
		return
	}

	product, err := prh.ProductRecServ.CreateProductRecords(productRecBody)
	if err != nil {
		if err, ok := err.(*appErr.GenericError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody(appErr.UnknowErr.Error(), nil))
		return
	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("success", product))
}

func (prh *ProductRecHandler) GetProductRecReport(w http.ResponseWriter, r *http.Request) {
	idProduct, err := strconv.Atoi(r.URL.Query().Get("id"))
	fmt.Println(idProduct)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("Invalid Parameter", nil))
		return
	}

	product, err := prh.ProductRecServ.GetProductRecordReport(idProduct)
	if err != nil {
		if err, ok := err.(*appErr.GenericError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody(appErr.UnknowErr.Error(), nil))
		return
	}


	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("success", product))
}