package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bootcamp-go/web/response"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler/responses"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	svc "github.com/maxwelbm/alkemy-g7.git/internal/service/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type CarrierHandler struct {
	srv svc.ICarrierService
}

func NewCarrierHandler(srv svc.ICarrierService) *CarrierHandler {
	return &CarrierHandler{srv: srv}
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
		var reqBody model.Carries

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			response.JSON(w, http.StatusBadRequest, responses.CreateResponseBody("invalid request body", nil))
			return
		}

		err := reqBody.ValidateEmptyFields(false)

		if err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, responses.CreateResponseBody(err.Error(), nil))
			return
		}

		carrier, err := h.srv.PostCarrier(reqBody)

		if err != nil {
			if err, ok := err.(*customerror.CarrierError); ok {
				response.JSON(w, err.Code, responses.CreateResponseBody(err.Error(), nil))
				return
			}

			if strings.Contains(err.Error(), customerror.ErrLocalityNotFound.Error()) {
				response.JSON(w, http.StatusNotFound, responses.CreateResponseBody(err.Error(), nil))
				return
			}

			response.JSON(w, http.StatusInternalServerError, responses.CreateResponseBody("unable to post carrier", nil))

			return
		}

		response.JSON(w, http.StatusCreated, responses.CreateResponseBody("", carrier))
	}
}
