package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customError"
)

type ProductBatchesJSON struct {
	ID                 int       `json:"id"`
	BatchNumber        string    `json:"batch_number"`
	CurrentQuantity    int       `json:"current_quantity"`
	CurrentTemperature float64   `json:"current_temperature"`
	MinimumTemperature float64   `json:"minimum_temperature"`
	DueDate            time.Time `json:"due_date"`
	InitialQuantity    int       `json:"initial_quantity"`
	ManufacturingDate  time.Time `json:"manufacturing_date"`
	ManufacturingHour  int       `json:"manufacturing_hour"`
	ProductID          int       `json:"product_id"`
	SectionID          int       `json:"section_id"`
}

type ProductBatchesController struct {
	sv *service.ProductBatchesService
}

func CreateProductBatchesHandler(sv *service.ProductBatchesService) *ProductBatchesController {
	return &ProductBatchesController{sv}
}

func (h *ProductBatchesController) Post(w http.ResponseWriter, r *http.Request) {
	var reqBody ProductBatchesJSON

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&reqBody)

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("invalid request body", nil))
		return
	}

	if reqBody == (ProductBatchesJSON{}) {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("request body cannot be empty", nil))
		return
	}

	productBatches := model.ProductBatches{
		BatchNumber:        reqBody.BatchNumber,
		CurrentQuantity:    reqBody.CurrentQuantity,
		CurrentTemperature: reqBody.CurrentTemperature,
		MinimumTemperature: reqBody.MinimumTemperature,
		DueDate:            reqBody.DueDate,
		InitialQuantity:    reqBody.InitialQuantity,
		ManufacturingDate:  reqBody.ManufacturingDate,
		ManufacturingHour:  reqBody.ManufacturingHour,
		ProductID:          reqBody.ProductID,
		SectionID:          reqBody.SectionID,
	}

	pb, err := h.sv.Post(&productBatches)

	if err != nil {
		if err, ok := err.(*customError.GenericError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			return
		}
		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to create product batches", nil))
		return
	}
	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", pb))
}
