package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	svc "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type CarrierHandler struct {
	Srv svc.ICarrierService
	log logger.Logger
}

func NewCarrierHandler(srv svc.ICarrierService, log logger.Logger) *CarrierHandler {
	return &CarrierHandler{Srv: srv, log: log}
}

// PostCarriers creates a new carrier.
// @Summary Create a new carrier
// @Description Creates a new carrier with the provided data
// @Tags Carriers
// @Accept json
// @Produce json
// @Param carrier body model.Carries true "Carrier data"
// @Success 201 {object} model.CarrierResponseSwagger
// @Failure 400 {object} model.ErrorResponseSwagger "Invalid request body"
// @Failure 422 {object} model.ErrorResponseSwagger "Invalid fields"
// @Failure 404 {object} model.ErrorResponseSwagger "Locality not found"
// @Failure 500 {object} model.ErrorResponseSwagger "Unable to post carrier"
// @Router /carriers [post]
func (h *CarrierHandler) PostCarriers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.log.Log("CarrierHandler", "INFO", "initializing PostCarriers function")
		var reqBody model.Carries

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			h.log.Log("CarrierHandler", "ERROR", "invalid request body")
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid request body", nil))
			return
		}

		err := reqBody.ValidateEmptyFields(false)

		if err != nil {
			h.log.Log("CarrierHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		carrier, err := h.Srv.PostCarrier(reqBody)

		if err != nil {
			if err, ok := err.(*customerror.CarrierError); ok {
				h.log.Log("CarrierHandler", "ERROR", fmt.Sprintf("Error: %v", err))
				response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
				return
			}

			if strings.Contains(err.Error(), customerror.ErrLocalityNotFound.Error()) {
				h.log.Log("CarrierHandler", "ERROR", fmt.Sprintf("Error: %v", err))
				response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
				return
			}

			h.log.Log("CarrierHandler", "ERROR", fmt.Sprintf("Error: %v", err))
			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to post carrier", nil))

			return
		}

		h.log.Log("CarrierHandler", "INFO", "PostCarriers completed successfully")
		response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", carrier))
	}
}
