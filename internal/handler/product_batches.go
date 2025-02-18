package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
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
	Sv  interfaces.IProductBatchesService
	log logger.Logger
}

func CreateProductBatchesHandler(sv interfaces.IProductBatchesService, log logger.Logger) *ProductBatchesController {
	return &ProductBatchesController{Sv: sv, log: log}
}

func (h *ProductBatchesController) Post(w http.ResponseWriter, r *http.Request) {
	h.log.Log("ProductBatchesController", "INFO", "initializing Post controller function")

	var reqBody ProductBatchesJSON

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&reqBody)

	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("invalid request body", nil))
		h.log.Log("ProductBatchesController", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	if reqBody == (ProductBatchesJSON{}) {
		response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody("request body cannot be empty", nil))
		h.log.Log("ProductBatchesController", "ERROR", fmt.Sprintf("Error: %v", err))

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

	pb, err := h.Sv.Post(&productBatches)

	if err != nil {
		if err, ok := err.(*customerror.GenericError); ok {
			response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
			h.log.Log("ProductBatchesController", "ERROR", fmt.Sprintf("Error: %v", err))

			return
		}

		response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("Unable to create product batches", nil))
		h.log.Log("ProductBatchesController", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", pb))
	h.log.Log("ProductBatchesController", "INFO", "post function executed successfully")
}
