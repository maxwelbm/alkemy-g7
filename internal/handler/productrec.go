package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"
	responses "github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
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
		response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
		return
	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("success", product))
}